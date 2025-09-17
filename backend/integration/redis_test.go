package integration

import (
	"context"
	"testing"

	"github.com/thomascpowell/drive/store"
	"github.com/thomascpowell/drive/utils"
)

func TestRedisConnection(t *testing.T) {
	rdb := store.NewRDB("localhost:6379")
	ctx := context.Background()
	pong, err := rdb.Client.Ping(ctx).Result()
	utils.Expect(t, err, nil, "unexpected error")
	utils.Expect(t, pong, "PONG", "unexpected redis response")
}
