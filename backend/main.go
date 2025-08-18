package main

import (
	"github.com/thomascpowell/drive/api"
	"github.com/thomascpowell/drive/jobs"
	"github.com/thomascpowell/drive/store"
	"log"
)

var QUEUE_SIZE = 10
var WORKER_COUNT = 4

func main() {
	store := store.NewStore("./data/app.db")
	dispatcher := jobs.NewDispatcher(store, QUEUE_SIZE)
	dispatcher.StartWorkers(WORKER_COUNT)
	router := api.SetupRouter(dispatcher)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("error starting gin: %v", err)
	}
}
