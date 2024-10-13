package cache

import (
	"NumismaticClubApi/pkg/api/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache[K, T any] struct {
	client   *redis.Client
	cacheKey string
	ttl      time.Duration
}

func NewRedisCache[K, T any](client *redis.Client, cacheKey string, ttl time.Duration) *RedisCache[K, T] {
	return &RedisCache[K, T]{
		client:   client,
		cacheKey: cacheKey,
		ttl:      ttl,
	}
}

func (r RedisCache[K, T]) Set(ctx utils.MyContext, key K, input T) {
	inputJSON, _ := json.Marshal(input)
	r.client.Set(ctx.Ctx, fmt.Sprintf(r.cacheKey, key), inputJSON, r.ttl)
}

func (r RedisCache[K, T]) Get(ctx utils.MyContext, key K) (T, error) {
	var result T

	bytes, err := r.client.Get(ctx.Ctx, fmt.Sprintf(r.cacheKey, key)).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return result, ErrNotFound
		}
		return result, err
	}

	err = json.Unmarshal(bytes, &result)

	return result, err
}

func (r RedisCache[K, T]) Delete(ctx utils.MyContext, key K) {
	r.client.Del(ctx.Ctx, fmt.Sprintf(r.cacheKey, key))
}
