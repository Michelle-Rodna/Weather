package cache

import (
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
	"context"
)
var client *redis.Client

func InitRedis(addr, password string, db int) error{
	client = redis.NewClient(&redis.Options{
		Addr: 	  addr,
		Password: password,
		DB: 	  db,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return client.Ping(ctx).Err()
}

func SetCache(ctx context.Context, key string, value any, ttl int) error{
	data, err := json.Marshal(value)
	if err != nil{
		return err
	}
	return client.Set(ctx, key, data, time.Duration(ttl)*time.Second).Err()
}

func GetCache(ctx context.Context, key string, box any) error{
	data, err := client.Get(ctx, key).Bytes()
	if err  != nil{
		return err6
	}
	return json.Unmarshal(data, box)
}

