package main

import (
	"log"

	"github.com/thomascpowell/drive/api"
	"github.com/thomascpowell/drive/jobs"
	"github.com/thomascpowell/drive/redis"
	"github.com/thomascpowell/drive/store"
	"github.com/thomascpowell/drive/utils"
)

var QUEUE_SIZE = 10
var WORKER_COUNT = 4

func main() {
	s := store.NewStore("./data/app.db")
	r := redis.NewRedis(utils.GetRedisURL())
	dispatcher := jobs.NewDispatcher(s, &r, QUEUE_SIZE)
	dispatcher.StartWorkers(WORKER_COUNT)
	router := api.SetupRouter(dispatcher)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("error starting gin: %v", err)
	}
}
