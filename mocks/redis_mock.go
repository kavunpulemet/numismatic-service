package mocks

import (
	"NumismaticClubApi/pkg/api/utils"
)

type RedisMock[K, T any] struct {
	SetCalls []struct {
		Ctx   utils.MyContext
		Key   K
		Input T
	}
	GetCalls []struct {
		Ctx utils.MyContext
		Key K
	}
	GetResults struct {
		Value T
		Err   error
	}
	DeleteCalls []struct {
		Ctx utils.MyContext
		Key K
	}
}

func (r *RedisMock[K, T]) Set(ctx utils.MyContext, key K, input T) {
	r.SetCalls = append(r.SetCalls, struct {
		Ctx   utils.MyContext
		Key   K
		Input T
	}{Ctx: ctx, Key: key, Input: input})
}

func (r *RedisMock[K, T]) Get(ctx utils.MyContext, key K) (T, error) {
	r.GetCalls = append(r.GetCalls, struct {
		Ctx utils.MyContext
		Key K
	}{Ctx: ctx, Key: key})

	return r.GetResults.Value, r.GetResults.Err
}

func (r *RedisMock[K, T]) Delete(ctx utils.MyContext, key K) {
	r.DeleteCalls = append(r.DeleteCalls, struct {
		Ctx utils.MyContext
		Key K
	}{Ctx: ctx, Key: key})
}
