package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
)

const (
	number = 160
	weight = 1
)

type Shards struct {
	name   string
	master *redis.Client
	slave  []*redis.Client
}

type ShardRedis struct {
	shard *TreeMap
}

func NewShard(name string, master *redis.Client, slave ...*redis.Client) *Shards {
	return &Shards{
		name:   name,
		master: master,
		slave:  slave,
	}
}

func NewShardRedis(arr ...*Shards) *ShardRedis {
	tree := &TreeMap{}
	for i, s := range arr {
		for j := 0; j < number; j++ {
			if s.name != "" {
				ts := fmt.Sprintf("%s*%d%d", s.name, weight, j)
				tree.Put(Hash(ts), s)
			} else {
				ts := fmt.Sprintf("SHARD-%d-NODE-%d", i, j)
				tree.Put(Hash(ts), s)
			}
		}
	}
	return &ShardRedis{tree}
}

func (sr *ShardRedis) Slave(key string) *redis.Client {
	slaves := sr.shard.FindCeiling(Hash(key)).V.(*Shards).slave
	if len(slaves) == 0{
		return nil
	}
	return slaves[rand.Intn(len(slaves))]
}

func (sr *ShardRedis) Master(key string) *redis.Client {
	return sr.shard.FindCeiling(Hash(key)).V.(*Shards).master
}
