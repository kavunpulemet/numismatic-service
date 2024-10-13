package mocks

import (
	"NumismaticClubApi/models"
	"NumismaticClubApi/pkg/api/utils"
	dbmodels "NumismaticClubApi/pkg/database/models"
)

type MongoMock struct {
	CreateCalls []struct {
		Ctx  utils.MyContext
		Coin models.Coin
	}
	CreateResults struct {
		ID  string
		Err error
	}

	GetAllCalls []struct {
		Ctx utils.MyContext
	}
	GetAllResults struct {
		Coins []models.Coin
		Err   error
	}

	GetByIdCalls []struct {
		Ctx    utils.MyContext
		CoinId string
	}
	GetByIdResults struct {
		Coin models.Coin
		Err  error
	}

	UpdateCalls []struct {
		Ctx    utils.MyContext
		CoinId string
		Input  dbmodels.UpdateCoin
	}
	UpdateResults struct {
		Err error
	}

	DeleteCalls []struct {
		Ctx    utils.MyContext
		CoinId string
	}
	DeleteResults struct {
		Err error
	}
}

func (m *MongoMock) Create(ctx utils.MyContext, coin models.Coin) (string, error) {
	m.CreateCalls = append(m.CreateCalls, struct {
		Ctx  utils.MyContext
		Coin models.Coin
	}{Ctx: ctx, Coin: coin})

	return m.CreateResults.ID, m.CreateResults.Err
}

func (m *MongoMock) GetAll(ctx utils.MyContext) ([]models.Coin, error) {
	m.GetAllCalls = append(m.GetAllCalls, struct {
		Ctx utils.MyContext
	}{Ctx: ctx})

	return m.GetAllResults.Coins, m.GetAllResults.Err
}

func (m *MongoMock) GetById(ctx utils.MyContext, coinId string) (models.Coin, error) {
	m.GetByIdCalls = append(m.GetByIdCalls, struct {
		Ctx    utils.MyContext
		CoinId string
	}{Ctx: ctx, CoinId: coinId})

	return m.GetByIdResults.Coin, m.GetByIdResults.Err
}

func (m *MongoMock) Update(ctx utils.MyContext, coinId string, input dbmodels.UpdateCoin) error {
	m.UpdateCalls = append(m.UpdateCalls, struct {
		Ctx    utils.MyContext
		CoinId string
		Input  dbmodels.UpdateCoin
	}{Ctx: ctx, CoinId: coinId, Input: input})

	return m.UpdateResults.Err
}

func (m *MongoMock) Delete(ctx utils.MyContext, coinId string) error {
	m.DeleteCalls = append(m.DeleteCalls, struct {
		Ctx    utils.MyContext
		CoinId string
	}{Ctx: ctx, CoinId: coinId})

	return m.DeleteResults.Err
}
