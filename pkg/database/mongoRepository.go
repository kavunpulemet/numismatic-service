package database

import (
	"NumismaticClubApi/models"
	"NumismaticClubApi/pkg/api/utils"
	dbmodels "NumismaticClubApi/pkg/database/models"
	_ "embed"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CoinRepository interface {
	Create(ctx utils.MyContext, coin models.Coin) (string, error)
	GetAll(ctx utils.MyContext) ([]models.Coin, error)
	GetById(ctx utils.MyContext, coinId string) (models.Coin, error)
	Update(ctx utils.MyContext, coinId string, input dbmodels.UpdateCoin) error
	Delete(ctx utils.MyContext, coinId string) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{collection: db.Collection("coins")}
}

func (r *MongoRepository) Create(ctx utils.MyContext, coin models.Coin) (string, error) {
	coin.Id = uuid.New().String()

	_, err := r.collection.InsertOne(ctx.Ctx, coin)
	if err != nil {
		return "", err
	}

	return coin.Id, nil
}

func (r *MongoRepository) GetAll(ctx utils.MyContext) ([]models.Coin, error) {
	var coins []models.Coin

	cursor, err := r.collection.Find(ctx.Ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx.Ctx)

	for cursor.Next(ctx.Ctx) {
		var coin models.Coin
		if err := cursor.Decode(&coin); err != nil {
			return nil, err
		}
		coins = append(coins, coin)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return coins, nil
}

func (r *MongoRepository) GetById(ctx utils.MyContext, coinId string) (models.Coin, error) {
	var coin models.Coin

	err := r.collection.FindOne(ctx.Ctx, bson.M{"_id": coinId}).Decode(&coin)
	if err != nil {
		return models.Coin{}, err
	}

	return coin, nil
}

func (r *MongoRepository) Update(ctx utils.MyContext, coinId string, input dbmodels.UpdateCoin) error {
	update := bson.M{"$set": input}
	result, err := r.collection.UpdateOne(ctx.Ctx, bson.M{"_id": coinId}, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no document found for update")
	}

	return nil
}

func (r *MongoRepository) Delete(ctx utils.MyContext, coinId string) error {
	result, err := r.collection.DeleteOne(ctx.Ctx, bson.M{"_id": coinId})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found for deletion")
	}

	return nil
}
