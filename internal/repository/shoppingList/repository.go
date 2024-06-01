package shoppingList

import (
	"context"

	"github.com/IgorCooli/minha-cozinha-be/internal/business/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	InsertOne(ctx context.Context, expense model.Item) error
	InsertMany(ctx context.Context, expenses []model.Item) error
	Search(ctx context.Context, name string) []model.Item
	RemoveItem(ctx context.Context, id string) error
}

func NewRepository(client *mongo.Client) Repository {
	return mongoRepository{
		stockDB: client.Database("TCCMongoDB").Collection("shopping-list"),
	}
}

type mongoRepository struct {
	stockDB *mongo.Collection
}

func (r mongoRepository) InsertOne(ctx context.Context, stockItem model.Item) error {

	_, err := r.stockDB.InsertOne(ctx, stockItem)

	if err != nil {
		panic("Could not insert item")
	}

	return nil
}

func (r mongoRepository) InsertMany(ctx context.Context, expenses []model.Item) error {

	var input []interface{}
	for _, exp := range expenses {
		input = append(input, exp)
	}

	_, err := r.stockDB.InsertMany(ctx, input)
	if err != nil {
		panic("Could not insert items")
	}

	return nil
}

func (r mongoRepository) Search(ctx context.Context, name string) []model.Item {
	var results []model.Item

	filter := bson.D{}

	if name != "" {
		filter = bson.D{
			{"name", bson.M{"$regex": primitive.Regex{Pattern: name, Options: "i"}}},
		}
	}

	cursor, err := r.stockDB.Find(ctx, filter)
	if err != nil {
		return []model.Item{}
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result model.Item
		if err := cursor.Decode(&result); err != nil {
			panic(err)
		}
		results = append(results, result)
	}

	return results
}

func (r mongoRepository) RemoveItem(ctx context.Context, id string) error {
	filter := bson.D{
		{"id", id},
	}

	_, err := r.stockDB.DeleteOne(ctx, filter)

	return err
}
