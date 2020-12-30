package redis

import (
	"bitbucket.org/fireflyon/go-library-template/pkg/cache"
	"context"
	rd "github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type redisConfig struct {
	address string
	db int
}

type (
	redisAdapter struct {
		redisConfig
		redisClient *rd.Client
	}
)


// New creates an instance of Redis cache driver
func NewRedisAdapter(addr string, db int) cache.Cache {
	config := redisConfig{address: addr, db: db}
	return &redisAdapter{config, nil}
}

// New redis client
func (ra *redisAdapter) createRedisClient() *rd.Client {
	ra.redisClient = rd.NewClient(&rd.Options{
		Addr:     ra.redisConfig.address,
		Password: "",
		DB:       ra.redisConfig.db,
	})
	return ra.redisClient
}

// get redis client
func (ra *redisAdapter) getRedisClient() *rd.Client {
	if ra.redisClient != nil{
		return ra.redisClient
	}
	return ra.createRedisClient()
}


// Contains checks if cached key exists in Redis storage
func (ra *redisAdapter) Contains(key string) (int64, error) {
	client := ra.getRedisClient()
	return client.Exists(ctx, key).Result()
}

// Delete the cached key from Redis storage
func (ra *redisAdapter) Delete(key string) error {
	client := ra.getRedisClient()
	return client.Del(ctx, key).Err()
}

// Get retrieves the cached value from key of the Redis storage
func (ra *redisAdapter) Get(key string) (string, error) {
	client := ra.getRedisClient()
	return client.Get(ctx, key).Result()
}

// HGet returns the value associated with field in the hash stored at key
func (ra *redisAdapter) HGet(key string, field string) (string, error) {
	client := ra.getRedisClient()
	return client.HGet(ctx, key, field).Result()
}

// HGetAll retrieves multiple cached value from keys of the Redis storage
func (ra *redisAdapter) HGetAll(key string) (map[string]string, error) {
	client := ra.getRedisClient()
	return client.HGetAll(ctx, key).Result()
}

// Set a value in Redis storage by key
func (ra *redisAdapter) Set(key string, value interface{}, lifeTime time.Duration) error {
	client := ra.getRedisClient()
	return client.Set(ctx, key, value, lifeTime).Err()
}

// HSet sets field in the hash stored at key to value
func (ra *redisAdapter) HSet(key string, values ...interface{}) error {
	client := ra.getRedisClient()
	return client.HSet(ctx, key, values).Err()
}
