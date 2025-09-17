package store

import (
	"github.com/redis/go-redis/v9"
)

type RDB struct {
	Client *redis.Client
}

func NewRDB(Addr string) RDB {
	rdb := redis.NewClient(&redis.Options{
		Addr: Addr,
	})

	return RDB {
		Client: rdb,
	}
}

// TODO: write methods for generating share urls
