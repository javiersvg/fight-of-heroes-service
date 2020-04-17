package aggregates

import (
	"log"

	"github.com/javiersvg/fight-of-heroes-service/events"
	"github.com/javiersvg/fight-of-heroes-service/uuid"
)

type Fights struct {
	bus   events.VisitableBus
	cache map[string]*Fight
	store *events.EventStore
}

func NewFights(bus events.VisitableBus, store *events.EventStore) *Fights {
	return &Fights{bus, make(map[string]*Fight), store}
}

func (f *Fights) Initialize() {
	values, err := f.store.GetActiveEvents()
	if err != nil {
		log.Fatal(err)
	}
	for _, id := range *values {
		if events, ok := f.store.GetEvents(id); ok {
			fight := NewFight(id)
			fight.Load(events)
			f.cache[id] = fight
		}
	}
}

func (f *Fights) CreateFight(heroes []string) {
	event := events.NewFightCreated(uuid.NewUUID(), heroes)
	f.bus.Publish(event)
}

func (f *Fights) GetFights() *[]*Fight {
	fights := []*Fight{}
	for _, v := range f.cache {
		fights = append(fights, v)
	}
	return &fights
}

func (f *Fights) GetFight(id string) *Fight {
	if events, ok := f.store.GetEvents(id); ok {
		fight := NewFight(id)
		fight.Load(events)
		f.cache[id] = fight
		return fight
	}
	return nil
}

func (f *Fights) HandleFightCreated(event *events.FightCreated) {
	fight := NewFight(event.GetId())
	f.cache[event.GetId()] = fight
	newEvent := events.NewHeroesUpdated(event.GetId(), event.GetHeroes())
	f.bus.Publish(newEvent)
}

func (f *Fights) HandleHeroesUpdated(event *events.HeroesUpdated) {
	f.cache[event.GetId()].HandleHeroesUpdated(event)
}

type FightsVisitor struct {
	handler *Fights
}

func NewFightsVisitor(fights *Fights) *FightsVisitor {
	return &FightsVisitor{fights}
}

func (v *FightsVisitor) VisitForFightCreated(event *events.FightCreated) {
	v.handler.HandleFightCreated(event)
}

func (v *FightsVisitor) VisitForHeroesUpdated(event *events.HeroesUpdated) {
	v.handler.HandleHeroesUpdated(event)
}

func (f *Fights) UpdateFights(id string, heroes []string) {
	event := events.NewHeroesUpdated(id, heroes)
	f.bus.Publish(event)
}

func (f *Fights) DeleteFight(id string) {
	event := events.NewFightDeleted(id)
	f.bus.Publish(event)
}

func (f *FightsVisitor) VisitForFightDeleted(event *events.FightDeleted) {
	f.handler.HandleFightDeleted(event)
}

func (f *Fights) HandleFightDeleted(event *events.FightDeleted) {
	delete(f.cache, event.GetId())
}
