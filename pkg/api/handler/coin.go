package handler

import (
	"NumismaticClubApi/models"
	"NumismaticClubApi/pkg/api/utils"
	"NumismaticClubApi/pkg/service/coin"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Create godoc
// @Summary Create a new coin
// @Description Adds a new coin to the collection
// @Tags coins
// @Accept  json
// @Produce  json
// @Param coin body models.Coin true "Coin data"
// @Success 200 {object} map[string]interface{} "id of the created coin"
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /coins/ [post]
func Create(ctx utils.MyContext, service coin.CoinService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var coin models.Coin
		if err := json.NewDecoder(r.Body).Decode(&coin); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		coinId, err := service.Create(coin)
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"id": coinId,
		}

		if err = utils.WriteResponse(w, http.StatusOK, response); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// GetAll godoc
// @Summary Get all coins
// @Description Retrieves all coins from the collection
// @Tags coins
// @Produce  json
// @Success 200 {array} models.Coin "List of coins"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /coins/ [get]
func GetAll(ctx utils.MyContext, service coin.CoinService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		coins, err := service.GetAll()
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = utils.WriteResponse(w, http.StatusOK, coins); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// GetById godoc
// @Summary Get coin by ID
// @Description Retrieves a coin by its ID
// @Tags coins
// @Produce  json
// @Param id path int true "Coin ID"
// @Success 200 {object} models.Coin "Coin data"
// @Failure 400 {object} utils.ErrorResponse "Invalid coin ID"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /coins/{id}/ [get]
func GetById(ctx utils.MyContext, service coin.CoinService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		coinId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		coins, err := service.GetById(coinId)
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = utils.WriteResponse(w, http.StatusOK, coins); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Update godoc
// @Summary Update coin
// @Description Updates an existing coin by its ID
// @Tags coins
// @Accept  json
// @Produce  json
// @Param id path int true "Coin ID"
// @Param coin body models.Coin true "Updated coin data"
// @Success 200 {object} utils.StatusResponse "ok"
// @Failure 400 {object} utils.ErrorResponse "Invalid coin ID or bad request"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /coins/{id}/ [put]
func Update(ctx utils.MyContext, service coin.CoinService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		coinId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		var input models.Coin
		if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = service.Update(coinId, input); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = utils.WriteResponse(w, http.StatusOK, utils.StatusResponse{Status: "ok"}); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Delete godoc
// @Summary Delete coin
// @Description Deletes a coin by its ID
// @Tags coins
// @Param id path int true "Coin ID"
// @Success 200 {object} utils.StatusResponse "ok"
// @Failure 400 {object} utils.ErrorResponse "Invalid coin ID"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /coins/{id}/ [delete]
func Delete(ctx utils.MyContext, service coin.CoinService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		coinId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = service.Delete(coinId); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = utils.WriteResponse(w, http.StatusOK, utils.StatusResponse{Status: "ok"}); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
