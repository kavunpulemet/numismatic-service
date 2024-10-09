package coin

import (
	"NumismaticClubApi/models"
	"NumismaticClubApi/pkg/api/utils"
	"NumismaticClubApi/pkg/repository"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

const CacheKeyAllCoins = "all_coins"

type CoinService interface {
	Create(ctx utils.MyContext, coin models.Coin) (string, error)
	GetAll(ctx utils.MyContext) ([]models.Coin, error)
	GetById(ctx utils.MyContext, coinId string) (models.Coin, error)
	Update(ctx utils.MyContext, coinId string, input models.Coin) error
	Delete(ctx utils.MyContext, coinId string) error
}

type ImplCoin struct {
	repo        repository.CoinRepository
	redisClient *redis.Client
}

func NewCoinService(repo repository.CoinRepository, client *redis.Client) *ImplCoin {
	return &ImplCoin{
		repo:        repo,
		redisClient: client,
	}
}

func (s *ImplCoin) Create(ctx utils.MyContext, coin models.Coin) (string, error) {
	coinId, err := s.repo.Create(ctx, coin)
	if err != nil {
		return coinId, err
	}

	s.redisClient.Del(ctx.Ctx, CacheKeyAllCoins)

	return coinId, nil
}

func (s *ImplCoin) GetAll(ctx utils.MyContext) ([]models.Coin, error) {
	var coins []models.Coin

	cachedCoins, err := s.redisClient.Get(ctx.Ctx, CacheKeyAllCoins).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(cachedCoins), &coins); err == nil {
			return coins, nil
		}
	}

	coins, err = s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	coinsJSON, _ := json.Marshal(coins)
	s.redisClient.Set(ctx.Ctx, CacheKeyAllCoins, coinsJSON, time.Minute*10)

	return coins, nil
}

func (s *ImplCoin) GetById(ctx utils.MyContext, coinId string) (models.Coin, error) {
	coin, err := s.repo.GetById(ctx, coinId)

	return coin, err
}

func (s *ImplCoin) Update(ctx utils.MyContext, coinId string, input models.Coin) error {
	err := s.repo.Update(ctx, coinId, input)
	if err != nil {
		return err
	}

	s.redisClient.Del(ctx.Ctx, CacheKeyAllCoins)

	return nil
}

func (s *ImplCoin) Delete(ctx utils.MyContext, coinId string) error {
	err := s.repo.Delete(ctx, coinId)
	if err != nil {
		return err
	}

	s.redisClient.Del(ctx.Ctx, CacheKeyAllCoins)

	return nil
}
