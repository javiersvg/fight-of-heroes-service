package events

type VisitableBus interface {
	Publish(Visitable)
	Subscribe(Visitor)
}

type FightOfHeroesBus struct {
	Subscribers []Visitor
}

func NewFightOfHeroesBus() *FightOfHeroesBus {
	return &FightOfHeroesBus{[]Visitor{}}
}

func (b *FightOfHeroesBus) Publish(event Visitable) {
	for _,subscriber := range b.Subscribers {
		go event.Accept(subscriber)
	}
}

func (b *FightOfHeroesBus) Subscribe(subscriber Visitor) {
	b.Subscribers = append(b.Subscribers, subscriber)
}