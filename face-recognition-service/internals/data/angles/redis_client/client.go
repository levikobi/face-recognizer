package redis_client

import (
	"context"
	"encoding/json"
	"face-recognition-service/internals/models"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type redisDB struct {
	config models.Config
	db *redis.Client
}

func Connect(config models.Config) *redisDB {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
	})

	log.Println("connected to redis")

	return &redisDB{config, rdb}
}

func (this *redisDB) Get(ctx context.Context, key string) (map[string]float64, bool, error) {
	cmd := this.db.Get(ctx, key)
	if cmd.Err() == redis.Nil {
		return nil, false, nil
	}

	cache := map[string]float64{}
	if err := json.Unmarshal([]byte(cmd.Val()), &cache); err != nil {
		return nil, false, err
	}
	return cache, true, nil
}

func (this *redisDB) Set(ctx context.Context, key string, value *map[string]float64) error {
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	status := this.db.Set(ctx, key, string(valueBytes), time.Duration(this.config.RedisTTLSeconds) * time.Second)
	return status.Err()
}
