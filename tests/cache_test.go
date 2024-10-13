package tests

import (
	"NumismaticClubApi/mocks"
	"NumismaticClubApi/models"
	"NumismaticClubApi/pkg/api/utils"
	"NumismaticClubApi/pkg/database"
	"NumismaticClubApi/pkg/database/cache"
	"NumismaticClubApi/pkg/service/coin"
	"context"
	"errors"
	"testing"
)

func TestGetById_FoundInCache_MongoNotCalled(t *testing.T) {
	ctx := utils.MyContext{Ctx: context.TODO(), Logger: nil}
	coinId := "1"
	expectedCoin := models.Coin{Id: "1", Name: "Nicecoinbro"}

	redisMock := &mocks.RedisMock[string, models.Coin]{}
	redisMock.GetResults = struct {
		Value models.Coin
		Err   error
	}{Value: expectedCoin, Err: nil}

	mongoMock := &mocks.MongoMock{}

	service := coin.NewCoinService(mongoMock, redisMock)

	coin, err := service.GetById(ctx, coinId)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
	if coin != expectedCoin {
		t.Errorf("Expected coin: %v, but got: %v", expectedCoin, coin)
	}

	if len(mongoMock.GetByIdCalls) != 0 {
		t.Error("Expected MongoDB GetById not to be called")
	}
}

func TestGetById_NotFoundInCache_FoundInMongo(t *testing.T) {
	ctx := utils.MyContext{Ctx: context.TODO(), Logger: nil}
	coinId := "1"
	expectedCoin := models.Coin{Id: "1", Name: "Nicecoinbro"}

	redisMock := &mocks.RedisMock[string, models.Coin]{}
	redisMock.GetResults = struct {
		Value models.Coin
		Err   error
	}{Value: models.Coin{}, Err: cache.ErrNotFound}

	mongoMock := &mocks.MongoMock{}
	mongoMock.GetByIdResults = struct {
		Coin models.Coin
		Err  error
	}{Coin: expectedCoin, Err: nil}

	service := coin.NewCoinService(mongoMock, redisMock)

	coin, err := service.GetById(ctx, coinId)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
	if coin != expectedCoin {
		t.Errorf("Expected coin: %v, but got: %v", expectedCoin, coin)
	}

	if len(mongoMock.GetByIdCalls) == 0 {
		t.Error("Expected MongoDB GetById to be called")
	}

	if len(redisMock.SetCalls) == 0 {
		t.Error("Expected redis.Set to be called")
	}
	if redisMock.SetCalls[0].Key != coinId {
		t.Errorf("Expected redis.Set to be called with key %s, but got %s", coinId, redisMock.SetCalls[0].Key)
	}

}

func TestGetById_NotFoundInCache_NotFoundInMongo(t *testing.T) {
	ctx := utils.MyContext{Ctx: context.TODO(), Logger: nil}
	coinId := "1"

	redisMock := &mocks.RedisMock[string, models.Coin]{}
	redisMock.GetResults = struct {
		Value models.Coin
		Err   error
	}{Value: models.Coin{}, Err: cache.ErrNotFound}

	mongoMock := &mocks.MongoMock{}
	mongoMock.GetByIdResults = struct {
		Coin models.Coin
		Err  error
	}{Coin: models.Coin{}, Err: database.ErrNotFound}

	service := coin.NewCoinService(mongoMock, redisMock)

	_, err := service.GetById(ctx, coinId)

	if err == nil {
		t.Fatal("Expected an error, but got nil")
	}
	if !errors.Is(err, database.ErrNotFound) {
		t.Errorf("Expected error: %v, but got: %v", database.ErrNotFound, err)
	}

	if len(redisMock.SetCalls) != 0 {
		t.Error("Expected redis.Set not to be called")
	}

	if len(mongoMock.GetByIdCalls) == 0 {
		t.Error("Expected MongoDB GetById to be called")
	}
}

func TestUpdate_DeletesFromCache(t *testing.T) {
	ctx := utils.MyContext{Ctx: context.TODO(), Logger: nil}
	coinId := "1"
	input := models.Coin{Name: "UpdatedCoin"}

	redisMock := &mocks.RedisMock[string, models.Coin]{}
	mongoMock := &mocks.MongoMock{}
	mongoMock.UpdateResults.Err = nil

	service := coin.NewCoinService(mongoMock, redisMock)

	err := service.Update(ctx, coinId, input)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(redisMock.DeleteCalls) == 0 {
		t.Errorf("Expected redis.Delete to be called")
	} else if redisMock.DeleteCalls[0].Key != coinId {
		t.Errorf("Expected redis.Delete to be called with key %s, but got %s", coinId, redisMock.DeleteCalls[0].Key)
	}
}

func TestDelete_DeletesFromCache(t *testing.T) {
	ctx := utils.MyContext{Ctx: context.TODO(), Logger: nil}
	coinId := "1"

	redisMock := &mocks.RedisMock[string, models.Coin]{}
	mongoMock := &mocks.MongoMock{}
	mongoMock.DeleteResults.Err = nil

	service := coin.NewCoinService(mongoMock, redisMock)

	err := service.Delete(ctx, coinId)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(redisMock.DeleteCalls) == 0 {
		t.Errorf("Expected redis.Delete to be called")
	} else if redisMock.DeleteCalls[0].Key != coinId {
		t.Errorf("Expected redis.Delete to be called with key %s, but got %s", coinId, redisMock.DeleteCalls[0].Key)
	}
}
