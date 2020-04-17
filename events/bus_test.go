package events

import (
	"testing"
)

type TestVisitor struct {
	fightCreated  *FightCreated
	heroesUpdated *HeroesUpdated
	fightDeleted  *FightDeleted
}

func (v *TestVisitor) VisitForFightCreated(e *FightCreated) {
	v.fightCreated = e
}
func (v *TestVisitor) VisitForHeroesUpdated(e *HeroesUpdated) {
	v.heroesUpdated = e
}
func (v *TestVisitor) VisitForFightDeleted(e *FightDeleted) {
	v.fightDeleted = e
}

func TestNewFightOfHeroesBus(t *testing.T) {
	b := NewFightOfHeroesBus()
	if b == nil {
		t.Error("NewFightOfHeroesBus() should not be null")
	}
}

func TestSubscribeAndPublish(t *testing.T) {
	b := NewFightOfHeroesBus()
	v := TestVisitor{}
	e := NewFightCreated("TEST_ID", make([]string, 0))
	b.Subscribe(&v)
	b.Publish(e)
	if e != v.fightCreated {
		t.Errorf("Visitor should have been called with fight created event %v != %v", *e, v.fightCreated)
	}
}
