package coin

import (
	"NumismaticClubApi/models"
	"NumismaticClubApi/pkg/repository"
)

type CoinService interface {
	Create(coin models.Coin) (int, error)
	GetAll() ([]models.Coin, error)
	GetById(coinId int) (models.Coin, error)
	Update(coinId int, input models.Coin) error
	Delete(coinId int) error
}

type ImplCoin struct {
	repo repository.CoinRepository
}

func NewCoinService(repo repository.CoinRepository) *ImplCoin {
	return &ImplCoin{repo: repo}
}

func (s *ImplCoin) Create(coin models.Coin) (int, error) {
	return s.repo.Create(coin)
}

func (s *ImplCoin) GetAll() ([]models.Coin, error) {
	coins, err := s.repo.GetAll()

	return coins, err
}

func (s *ImplCoin) GetById(coinId int) (models.Coin, error) {
	coin, err := s.repo.GetById(coinId)

	return coin, err
}

func (s *ImplCoin) Update(coinId int, input models.Coin) error {
	return s.repo.Update(coinId, input)
}

func (s *ImplCoin) Delete(coinId int) error {
	return s.repo.Delete(coinId)
}
