package core_redis

import (
	core_config "auth/internal/core/config"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	rdb *redis.Client
}

func NewClient(
	cfg *core_config.Config,
) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB: 0,
	})
}

func NewRedisStorage(
	rdb *redis.Client,
) *RedisStorage {
	return &RedisStorage{
		rdb: rdb,
	}
}

func (r *RedisStorage) Set(
	ctx context.Context,
	key string,
	val interface{},
	ttl time.Duration,
) error {
	return 	r.rdb.Set(ctx, key, val, ttl).Err()
}

func (r *RedisStorage) Get(
	ctx context.Context,
	key string,
) (string, error) {
	return r.rdb.Get(ctx, key).Result()
}

func (r *RedisStorage) Del(
	ctx context.Context,
	key string,
) error {
	return r.rdb.Del(ctx, key).Err()
}