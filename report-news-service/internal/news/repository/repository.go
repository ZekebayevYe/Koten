package repository

import (
	"context"

	"reportnewsservice/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewsRepository interface {
	Insert(ctx context.Context, news *proto.News) error
	FetchAll(ctx context.Context) ([]*proto.News, error)
}

type newsRepository struct {
	collection *mongo.Collection
}

func NewNewsRepository(db *mongo.Database) NewsRepository {
	return &newsRepository{
		collection: db.Collection("news"),
	}
}

func (r *newsRepository) Insert(ctx context.Context, news *proto.News) error {
	_, err := r.collection.InsertOne(ctx, news)
	return err
}

func (r *newsRepository) FetchAll(ctx context.Context) ([]*proto.News, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var newsList []*proto.News
	for cursor.Next(ctx) {
		var n proto.News
		if err := cursor.Decode(&n); err != nil {
			return nil, err
		}
		newsList = append(newsList, &n)
	}
	return newsList, nil
}
