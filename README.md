# go-shardJedis

> a implement of ShardJedis base on "github.com/go-redis/redis"

## Install

go get "github.com/zyfcn/go-shardJedis"

## Example

### Create a ShardInfo
`
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
`

### Create a ShardJedis
`
jedis := NewShardRedis(info)
`

### Use ShardJedis
`
client := jedis.Master("abc")
`

### Then you can do some with the client from github.com/go-redis/redis


