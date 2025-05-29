package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Document struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"user_id"`
	Filename  string             `bson:"filename"`
	Type      string             `bson:"type"`
	Content   []byte             `bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
}

func InsertDocument(ctx context.Context, coll *mongo.Collection, doc Document) (primitive.ObjectID, error) {
	res, err := coll.InsertOne(ctx, doc)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func GetDocumentsByUserID(ctx context.Context, coll *mongo.Collection, userID string) ([]Document, error) {
	var results []Document
	cursor, err := coll.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func GetDocumentByID(ctx context.Context, coll *mongo.Collection, id string) (*Document, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var doc Document
	err = coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&doc)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

func DeleteDocumentByID(ctx context.Context, coll *mongo.Collection, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = coll.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}
