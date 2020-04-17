package events

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/javiersvg/fight-of-heroes-service/clients"
)

type EventStore struct {
	db     *sql.DB
	mapper func([]*string) Visitable
}

type EventStoreVisitor struct {
	handler *EventStore
}

func NewEventStore() *EventStore {
	return &EventStore{clients.MySqlDatabaseFactory(), EventMapper()}
}

func NewEventStoreVisitor(eventStore *EventStore) *EventStoreVisitor {
	return &EventStoreVisitor{eventStore}
}

func (e *EventStore) GetEvents(id string) ([]Visitable, bool) {
	var name string
	var aggregateId string
	var data string
	stmt, err := e.db.Prepare("SELECT * FROM EVENTS WHERE AGGREGATE_ID = ?")
	defer stmt.Close()
	if err != nil {
		return nil, false
	}
	rows, err := stmt.Query(id)
	value := make([]Visitable, 0)
	for rows.Next() {
		if err := rows.Scan(&name, &aggregateId, &data); err != nil {
			log.Fatal(err)
		}
		value = append(value, e.mapper([]*string{&name, &aggregateId, &data}))
	}
	return value, len(value) > 0
}

func (e *EventStore) GetActiveEvents() (*[]string, error) {
	stmt, err := e.db.Prepare("SELECT DISTINCT AGGREGATE_ID FROM EVENTS WHERE AGGREGATE_ID NOT IN (SELECT AGGREGATE_ID FROM EVENTS WHERE EVENT_TYPE = 'FightDeleted')")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	values := make([]string, 0)
	var id string
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		values = append(values, id)
	}
	return &values, nil
}

func (e *EventStore) parseRows(rows *sql.Rows, c chan string) {
	var id string
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		c <- id
	}
	close(c)
}

func (v *EventStoreVisitor) VisitForFightCreated(event *FightCreated) {
	v.handler.HandleFightCreated(event)
}

func (e *EventStore) HandleFightCreated(event *FightCreated) {
	stmt, err := e.db.Prepare("INSERT INTO EVENTS VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmt.Close()
	value, _ := json.Marshal(event.GetHeroes())
	stmt.Exec("FightCreated", event.GetId(), value)
}

func (v *EventStoreVisitor) VisitForHeroesUpdated(event *HeroesUpdated) {
	v.handler.HandleHeroesUpdated(event)
}

func (e *EventStore) HandleHeroesUpdated(event *HeroesUpdated) {
	stmt, err := e.db.Prepare("INSERT INTO EVENTS VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmt.Close()
	value, _ := json.Marshal(event.GetHeroes())
	stmt.Exec("HeroesUpdated", event.GetId(), value)
}

func (e *EventStoreVisitor) VisitForFightDeleted(event *FightDeleted) {
	e.handler.HandleFightDeleted(event)
}

func (e *EventStore) HandleFightDeleted(event *FightDeleted) {
	stmt, err := e.db.Prepare("INSERT INTO EVENTS VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmt.Close()
	stmt.Exec("FightDeleted", event.GetId(), "")
}
