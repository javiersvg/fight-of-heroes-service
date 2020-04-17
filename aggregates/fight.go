package aggregates

import (
	"github.com/javiersvg/fight-of-heroes-service/events"
)

type Fight struct {
	Id     string
	Heroes []string
}

type FightVisitor struct {
	handler *Fight
}

func NewFight(id string) *Fight {
	return &Fight{id, []string{}}
}

func (f *Fight) Load(events []events.Visitable) {
	visitor := &FightVisitor{f}
	for _, event := range events {
		event.Accept(visitor)
	}
}

func (v *FightVisitor) VisitForFightCreated(event *events.FightCreated) {
	v.handler.HandleFightCreated(event)
}

func (f *Fight) HandleFightCreated(event *events.FightCreated) {
	//Do nothing
}

func (v *FightVisitor) VisitForHeroesUpdated(event *events.HeroesUpdated) {
	v.handler.HandleHeroesUpdated(event)
}

func (f *Fight) HandleHeroesUpdated(event *events.HeroesUpdated) {
	f.Heroes = event.GetHeroes()
}

func (v *FightVisitor) VisitForFightDeleted(event *events.FightDeleted) {
	//Do nothing
}
