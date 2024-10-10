package cache

import (
	"NumismaticClubApi/models"
	"NumismaticClubApi/pkg/api/utils"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache struct {
	client   *redis.Client
	cacheKey string
	ttl      time.Duration
}

func NewRedisCache(client *redis.Client, cacheKey string, ttl time.Duration) RedisCache {
	return RedisCache{
		client:   client,
		cacheKey: cacheKey,
		ttl:      ttl,
	}
}

func (r RedisCache) Set(ctx utils.MyContext, coins []models.Coin) {
	coinsJSON, _ := json.Marshal(coins)
	r.client.Set(ctx.Ctx, r.cacheKey, coinsJSON, r.ttl)
}

func (r RedisCache) Get(ctx utils.MyContext) ([]models.Coin, error) {
	var coins []models.Coin

	cachedCoins, err := r.client.Get(ctx.Ctx, r.cacheKey).Result()
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(cachedCoins), &coins); err == nil {
		return coins, nil
	}

	return nil, err
}

func (r RedisCache) Delete(ctx utils.MyContext) {
	r.client.Del(ctx.Ctx, r.cacheKey)
}
