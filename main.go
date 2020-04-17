package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/javiersvg/fight-of-heroes-service/aggregates"
	"github.com/javiersvg/fight-of-heroes-service/events"
	"github.com/javiersvg/fight-of-heroes-service/resources"
)

var fights *aggregates.Fights
var bus *events.FightOfHeroesBus
var store *events.EventStore

func init() {
	bus = events.NewFightOfHeroesBus()
	store = events.NewEventStore()
	fights = aggregates.NewFights(bus, store)
	fights.Initialize()
	bus.Subscribe(aggregates.NewFightsVisitor(fights))
	bus.Subscribe(events.NewEventStoreVisitor(store))
}

func main() {
	mux := setupHandlers()

	log.Println("Starting App...")
	if err := http.ListenAndServe(":8088", mux); err != nil {
		log.Fatal(err)
	}
}

func setupHandlers() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/command", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(resources.CommandResource(fights))))
	mux.Handle("/query/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(resources.QueriesResource(fights))))
	mux.Handle("/health", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(resources.HealthResource())))

	return mux
}
