package coin

import (
	"NumismaticClubApi/models"
	"NumismaticClubApi/pkg/api/utils"
	"NumismaticClubApi/pkg/repository"
)

type CoinService interface {
	Create(ctx utils.MyContext, coin models.Coin) (string, error)
	GetAll(ctx utils.MyContext) ([]models.Coin, error)
	GetById(ctx utils.MyContext, coinId string) (models.Coin, error)
	Update(ctx utils.MyContext, coinId string, input models.Coin) error
	Delete(ctx utils.MyContext, coinId string) error
}

type ImplCoin struct {
	repo repository.CoinRepository
}

func NewCoinService(repo repository.CoinRepository) *ImplCoin {
	return &ImplCoin{repo: repo}
}

func (s *ImplCoin) Create(ctx utils.MyContext, coin models.Coin) (string, error) {
	return s.repo.Create(ctx, coin)
}

func (s *ImplCoin) GetAll(ctx utils.MyContext) ([]models.Coin, error) {
	coins, err := s.repo.GetAll(ctx)

	return coins, err
}

func (s *ImplCoin) GetById(ctx utils.MyContext, coinId string) (models.Coin, error) {
	coin, err := s.repo.GetById(ctx, coinId)

	return coin, err
}

func (s *ImplCoin) Update(ctx utils.MyContext, coinId string, input models.Coin) error {
	return s.repo.Update(ctx, coinId, input)
}

func (s *ImplCoin) Delete(ctx utils.MyContext, coinId string) error {
	return s.repo.Delete(ctx, coinId)
}
