package cache

import (
	"NumismaticClubApi/pkg/api/utils"
)

type CoinCache[T any] interface {
	Set(ctx utils.MyContext, data []T)
	Get(ctx utils.MyContext) ([]T, error)
	Delete(ctx utils.MyContext)
}
