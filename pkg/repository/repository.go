package repository

import (
	"NumismaticClubApi/models"
	_ "embed"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type CoinRepository interface {
	Create(coin models.Coin) (int, error)
	GetAll() ([]models.Coin, error)
	GetById(coinId int) (models.Coin, error)
	Update(coinId int, input models.Coin) error
	Delete(coinId int) error
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

//go:embed sql/CreateCoin.sql
var createCoin string

func (r *Repository) Create(coin models.Coin) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	var id int
	err = tx.Get(&id, createCoin, coin.Name, coin.Country, coin.Year, coin.Denomination, coin.Material,
		coin.Weight, coin.Diameter, coin.Thickness, coin.Condition, coin.MintMark, coin.HistoricalInfo, coin.Value)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

//go:embed sql/GetAll.sql
var getAll string

func (r *Repository) GetAll() ([]models.Coin, error) {

	var coins []models.Coin

	err := r.db.Select(&coins, getAll)

	return coins, err
}

//go:embed sql/GetById.sql
var getById string

func (r *Repository) GetById(coinId int) (models.Coin, error) {
	var coin models.Coin

	err := r.db.Select(&coin, getById, coinId)

	return coin, err
}

func (r *Repository) Update(coinId int, input models.Coin) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var (
		coinUpdates []string
		args        []interface{}
		argId       = 1
	)

	if input.Name != "" {
		coinUpdates = append(coinUpdates, fmt.Sprintf("name = $%d", argId))
		args = append(args, input.Name)
		argId++
	}
	if input.Country != "" {
		coinUpdates = append(coinUpdates, fmt.Sprintf("country = $%d", argId))
		args = append(args, input.Country)
		argId++
	}
	if input.Year != 0 {
		coinUpdates = append(coinUpdates, fmt.Sprintf("year = $%d", argId))
		args = append(args, input.Year)
		argId++
	}
	if input.Denomination != "" {
		coinUpdates = append(coinUpdates, fmt.Sprintf("denomination = $%d", argId))
		args = append(args, input.Denomination)
		argId++
	}
	if input.Material != "" {
		coinUpdates = append(coinUpdates, fmt.Sprintf("material = $%d", argId))
		args = append(args, input.Material)
		argId++
	}
	if input.Weight != 0 {
		coinUpdates = append(coinUpdates, fmt.Sprintf("weight = $%d", argId))
		args = append(args, input.Weight)
		argId++
	}
	if input.Diameter != 0 {
		coinUpdates = append(coinUpdates, fmt.Sprintf("diameter = $%d", argId))
		args = append(args, input.Diameter)
		argId++
	}
	if input.Thickness != 0 {
		coinUpdates = append(coinUpdates, fmt.Sprintf("thickness = $%d", argId))
		args = append(args, input.Thickness)
		argId++
	}
	if input.Condition != "" {
		coinUpdates = append(coinUpdates, fmt.Sprintf("condition = $%d", argId))
		args = append(args, input.Condition)
		argId++
	}
	if input.MintMark != "" {
		coinUpdates = append(coinUpdates, fmt.Sprintf("mintmark = $%d", argId))
		args = append(args, input.MintMark)
		argId++
	}
	if input.HistoricalInfo != "" {
		coinUpdates = append(coinUpdates, fmt.Sprintf("historicalinfo = $%d", argId))
		args = append(args, input.HistoricalInfo)
		argId++
	}
	if input.Value != 0 {
		coinUpdates = append(coinUpdates, fmt.Sprintf("value = $%d", argId))
		args = append(args, input.Value)
		argId++
	}

	if len(coinUpdates) > 0 {
		query := fmt.Sprintf("UPDATE coins SET %s WHERE id = $%d",
			strings.Join(coinUpdates, ", "), argId)
		args = append(args, coinId)
		if _, err = tx.Exec(query, args...); err != nil {
			return err
		}
	}

	return nil
}

//go:embed sql/DeleteCoin.sql
var deleteCoin string

func (r *Repository) Delete(coinId int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(deleteCoin, coinId)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
