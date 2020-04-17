package events

import (
	"encoding/json"
)

type FightCreated struct {
	id     string
	heroes []string
}

type HeroesUpdated struct {
	id     string
	heroes []string
}

type FightDeleted struct {
	id string
}

type Visitable interface {
	Accept(Visitor)
}

type Visitor interface {
	VisitForFightCreated(*FightCreated)
	VisitForHeroesUpdated(*HeroesUpdated)
	VisitForFightDeleted(*FightDeleted)
}

func NewFightCreated(id string, heroes []string) *FightCreated {
	return &FightCreated{id, heroes}
}

func NewHeroesUpdated(id string, heroes []string) *HeroesUpdated {
	return &HeroesUpdated{id, heroes}
}

func NewFightDeleted(id string) *FightDeleted {
	return &FightDeleted{id}
}

func (f *FightCreated) GetId() string {
	return f.id
}

func (f *FightCreated) GetHeroes() []string {
	return f.heroes
}

func (h *HeroesUpdated) GetId() string {
	return h.id
}

func (f *FightDeleted) GetId() string {
	return f.id
}

func (h *HeroesUpdated) GetHeroes() []string {
	return h.heroes
}

func (f *FightCreated) Accept(visitor Visitor) {
	visitor.VisitForFightCreated(f)
}

func (h *HeroesUpdated) Accept(visitor Visitor) {
	visitor.VisitForHeroesUpdated(h)
}

func (f *FightDeleted) Accept(visitor Visitor) {
	visitor.VisitForFightDeleted(f)
}

func EventMapper() func([]*string) Visitable {
	return func(columns []*string) Visitable {
		switch *columns[0] {
		case "FightCreated":
			heroes := make([]string, 2)
			json.Unmarshal([]byte(*columns[2]), &heroes)
			return NewFightCreated(*columns[1], heroes)
		case "HeroesUpdated":
			heroes := make([]string, 2)
			json.Unmarshal([]byte(*columns[2]), &heroes)
			return NewHeroesUpdated(*columns[1], heroes)
		case "FightDeleted":
			return NewFightDeleted(*columns[1])
		default:
			panic("Unable to parse event: " + *columns[0])
		}
	}
}
