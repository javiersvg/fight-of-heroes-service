package events

import "testing"

func TestNewFightCreated(t *testing.T) {
	e := NewFightCreated("TEST_ID", []string{"TEST_HERO_1", "TEST_HERO_2"})

	if e == nil {
		t.Error("Event should not be nil")
	}
	if e.GetId() != "TEST_ID" {
		t.Errorf("Event should have id == TEST_ID instead of : %v", e.id)
	}
	if e.GetHeroes()[0] != "TEST_HERO_1" {
		t.Errorf("Event should have hero == TEST_HERO_1 instead of : %v", e.id)
	}
	if e.GetHeroes()[1] != "TEST_HERO_2" {
		t.Errorf("Event should have hero == TEST_HERO_2 instead of : %v", e.id)
	}
}

func TestNewHeroesUpdated(t *testing.T) {
	e := NewHeroesUpdated("TEST_ID", []string{"TEST_HERO_1", "TEST_HERO_2"})

	if e == nil {
		t.Error("Event should not be nil")
	}
	if e.GetId() != "TEST_ID" {
		t.Errorf("Event should have id == TEST_ID instead of : %v", e.id)
	}
	if e.GetHeroes()[0] != "TEST_HERO_1" {
		t.Errorf("Event should have hero == TEST_HERO_1 instead of : %v", e.id)
	}
	if e.GetHeroes()[1] != "TEST_HERO_2" {
		t.Errorf("Event should have hero == TEST_HERO_2 instead of : %v", e.id)
	}
}

func TestNewFightDeleted(t *testing.T) {
	e := NewFightDeleted("TEST_ID")

	if e == nil {
		t.Error("Event should not be nil")
	}
	if e.GetId() != "TEST_ID" {
		t.Errorf("Event should have id == TEST_ID instead of : %v", e.id)
	}
}

func TestFightCreatedAccept(t *testing.T) {
	e := NewFightCreated("TEST_ID", []string{"TEST_HERO_1", "TEST_HERO_2"})
	v := TestVisitor{}
	e.Accept(&v)

	if e != v.fightCreated {
		t.Errorf("Visitor should have been called with fight created event %v != %v", *e, v.fightCreated)
	}
}

func TestHeroesUpdated(t *testing.T) {
	e := NewHeroesUpdated("TEST_ID", []string{"TEST_HERO_1", "TEST_HERO_2"})
	v := TestVisitor{}
	e.Accept(&v)

	if e != v.heroesUpdated {
		t.Errorf("Visitor should have been called with fight created event %v != %v", *e, v.fightCreated)
	}
}

func TestFightDeleted(t *testing.T) {
	e := NewFightDeleted("TEST_ID")
	v := TestVisitor{}
	e.Accept(&v)

	if e != v.fightDeleted {
		t.Errorf("Visitor should have been called with fight created event %v != %v", *e, v.fightCreated)
	}
}
