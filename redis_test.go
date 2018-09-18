package redis

import (
	"testing"
	"fmt"
	"github.com/go-redis/redis"
)

func TestRedisClient(t *testing.T) {
	master := redis.NewClient(&redis.Options{
		Addr:     "localhost:6370",
		Password: "",
		DB:       0,
	})
	slave := redis.NewClient(&redis.Options{
		Addr:     "localhost:6371",
		Password: "",
		DB:       0,
	})
	info := NewShard("", master, slave)
	jedis := NewShardRedis(info)
	client := jedis.Master("abc")
	fmt.Println(client.Info().Result())
}
