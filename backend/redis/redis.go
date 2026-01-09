package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisInterface interface {
	Set(key string, value string) error
	Setex(key string, value string, ttl int) error
	Get(key string) (string, error)
	TTL(key string) (string, error)
}

var _ RedisInterface = (*Redis)(nil)

type Redis struct {
	Client *redis.Client
}

func NewRedis(Addr string) Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr: Addr,
	})

	return Redis{
		Client: rdb,
	}
}

func (rc *Redis) Set(key string, value string) error {
	ctx := context.Background()
	_, err := rc.Client.Set(ctx, key, value, 0).Result()
	return err
}

func (rc *Redis) Setex(key string, value string, ttl int) error {
	ctx, cancel := getCTX(2)
	defer cancel()
	// very stupid time.Duration requirement
	_, err := rc.Client.SetEx(ctx, key, value, time.Duration(ttl)*time.Second).Result()
	return err
}

func (rc *Redis) Get(key string) (string, error) {
	ctx := context.Background()
	return rc.Client.Get(ctx, key).Result()
}

func (rc *Redis) TTL(key string) (string, error) {
	ctx := context.Background()
	dur, err := rc.Client.TTL(ctx, key).Result()
	return dur.String(), err
}

func getCTX(expire int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(expire)*time.Second)
}
