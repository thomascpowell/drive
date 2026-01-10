package integration

import (
	"context"
	"github.com/thomascpowell/drive/redis"
	"github.com/thomascpowell/drive/utils"
	"net"
	"testing"
	"time"
)

func TestRedisConnection(t *testing.T) {
	timeout := 1 * time.Second
	conn, err := net.DialTimeout("tcp", "127.0.0.1:6379", timeout)
	if err != nil {
		t.Skip("Redis unavailable.")
		return
	}
	conn.Close()

	rdb := redis.NewRedis(utils.GetRedisURL())

	// test general function
	ctx := context.Background()
	pong, err := rdb.Client.Ping(ctx).Result()
	utils.Expect(t, err, nil, "unexpected error")
	utils.Expect(t, pong, "PONG", "unexpected redis response")

	// test wrappers
	rdb.Set("a", "some value")
	get, err := rdb.Get("a")
	utils.Expect(t, err, nil, "unexpected error")
	utils.Expect(t, get, "some value", "unexpected redis response")

	// test ttl
	rdb.Setex("b", "some other value", 1)
	// waste exactly 1001 milliseconds of my life
	time.Sleep(1001 * time.Millisecond)
	_, err = rdb.Get("b")
	utils.Expect(t, err.Error(), "redis: nil", "unexpected error")

	// do somthing unreasonable
	_, err = rdb.Get("some key that doesnt exist")
	utils.Expect(t, err.Error(), "redis: nil", "unexpected error")

	// make sure that connection is still fine
	rdb.Set("h", "h")
	get, err = rdb.Get("h")
	utils.Expect(t, get, "h", "unexpected redis response")
}
