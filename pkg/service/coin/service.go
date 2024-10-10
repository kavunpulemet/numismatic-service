package coin

import (
	"NumismaticClubApi/models"
	"NumismaticClubApi/pkg/api/utils"
	"NumismaticClubApi/pkg/database"
	"NumismaticClubApi/pkg/database/cache"
	"NumismaticClubApi/pkg/service/mappers"
)

type CoinService interface {
	Create(ctx utils.MyContext, coin models.Coin) (string, error)
	GetAll(ctx utils.MyContext) ([]models.Coin, error)
	GetById(ctx utils.MyContext, coinId string) (models.Coin, error)
	Update(ctx utils.MyContext, coinId string, input models.Coin) error
	Delete(ctx utils.MyContext, coinId string) error
}

type ImplCoinService struct {
	mongo database.CoinRepository
	redis cache.RedisCache
}

func NewCoinService(repo database.CoinRepository, cache cache.RedisCache) *ImplCoinService {
	return &ImplCoinService{
		mongo: repo,
		redis: cache,
	}
}

func (s *ImplCoinService) Create(ctx utils.MyContext, coin models.Coin) (string, error) {
	coinId, err := s.mongo.Create(ctx, coin)
	if err != nil {
		return coinId, err
	}

	s.redis.Delete(ctx)
	return coinId, nil
}

func (s *ImplCoinService) GetAll(ctx utils.MyContext) ([]models.Coin, error) {
	coins, err := s.redis.Get(ctx)
	if err == nil {
		return coins, nil
	}

	coins, err = s.mongo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	s.redis.Set(ctx, coins)

	return coins, nil
}

func (s *ImplCoinService) GetById(ctx utils.MyContext, coinId string) (models.Coin, error) {
	coin, err := s.mongo.GetById(ctx, coinId)

	return coin, err
}

func (s *ImplCoinService) Update(ctx utils.MyContext, coinId string, input models.Coin) error {
	err := s.mongo.Update(ctx, coinId, mappers.MapToUpdateCoin(input))
	if err != nil {
		return err
	}

	s.redis.Delete(ctx)

	return nil
}

func (s *ImplCoinService) Delete(ctx utils.MyContext, coinId string) error {
	err := s.mongo.Delete(ctx, coinId)
	if err != nil {
		return err
	}

	s.redis.Delete(ctx)

	return nil
}
