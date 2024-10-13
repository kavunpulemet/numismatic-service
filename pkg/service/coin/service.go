package coin

import (
	"NumismaticClubApi/models"
	"NumismaticClubApi/pkg/api/utils"
	"NumismaticClubApi/pkg/database"
	"NumismaticClubApi/pkg/database/cache"
	"NumismaticClubApi/pkg/service/mappers"
	"errors"
)

type CoinService interface {
	Create(ctx utils.MyContext, coin models.Coin) (string, error)
	GetAll(ctx utils.MyContext) ([]models.Coin, error)
	GetById(ctx utils.MyContext, coinId string) (models.Coin, error)
	Update(ctx utils.MyContext, coinId string, input models.Coin) error
	Delete(ctx utils.MyContext, coinId string) error
}

type ImplCoinService[K any, T any] struct {
	mongo database.CoinRepository
	cache cache.CoinCache[string, models.Coin] // *cache.RedisCache[string, models.Coin]
}

func NewCoinService(repo database.CoinRepository, cache cache.CoinCache[string, models.Coin]) *ImplCoinService[string, models.Coin] {
	return &ImplCoinService[string, models.Coin]{
		mongo: repo,
		cache: cache,
	}
}

func (s *ImplCoinService[K, T]) Create(ctx utils.MyContext, coin models.Coin) (string, error) {
	coinId, err := s.mongo.Create(ctx, coin)
	if err != nil {
		return coinId, err
	}

	return coinId, nil
}

func (s *ImplCoinService[K, T]) GetAll(ctx utils.MyContext) ([]models.Coin, error) {
	coins, err := s.mongo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return coins, nil
}

func (s *ImplCoinService[K, T]) GetById(ctx utils.MyContext, coinId string) (models.Coin, error) {
	coin, err := s.cache.Get(ctx, coinId)
	if err == nil {
		return coin, nil
	}

	if err != nil && !errors.Is(err, cache.ErrNotFound) {
		return coin, err
	}

	coin, err = s.mongo.GetById(ctx, coinId)
	if err != nil {
		return coin, err
	}
	s.cache.Set(ctx, coinId, coin)

	return coin, nil
}

func (s *ImplCoinService[K, T]) Update(ctx utils.MyContext, coinId string, input models.Coin) error {
	err := s.mongo.Update(ctx, coinId, mappers.MapToUpdateCoin(input))
	if err != nil {
		return err
	}

	s.cache.Delete(ctx, coinId)

	return nil
}

func (s *ImplCoinService[K, T]) Delete(ctx utils.MyContext, coinId string) error {
	err := s.mongo.Delete(ctx, coinId)
	if err != nil {
		return err
	}

	s.cache.Delete(ctx, coinId)

	return nil
}
