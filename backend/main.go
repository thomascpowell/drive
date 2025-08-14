package main

import (
	"github.com/thomascpowell/drive/api"
	"github.com/thomascpowell/drive/jobs"
	"github.com/thomascpowell/drive/store"
	"log"
)

func main() {
	store := store.NewStore("./data/app.db")
	dispatcher := jobs.NewDispatcher(store, 10)
	router := api.SetupRouter(dispatcher)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("error starting gin: %v", err)
	}
}
