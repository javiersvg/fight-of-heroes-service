package events

import "testing"

func TestNewFightOfHeroesBus(t *testing.T) {
	b := NewFightOfHeroesBus()
	if b == nil {
		t.Error("NewFightOfHeroesBus() should not be null")
	}
}
