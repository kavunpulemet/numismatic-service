package cache

import (
	"NumismaticClubApi/pkg/api/utils"
	"errors"
)

type CoinCache[K, T any] interface {
	Set(ctx utils.MyContext, key K, input T)
	Get(ctx utils.MyContext, key K) (T, error)
	Delete(ctx utils.MyContext, key K)
}

var ErrNotFound = errors.New("cache: item not found")
