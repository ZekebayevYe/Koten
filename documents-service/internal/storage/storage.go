package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage interface {
	InsertDocument(ctx context.Context, coll *mongo.Collection, doc Document) (primitive.ObjectID, error)
	GetDocumentsByUserID(ctx context.Context, coll *mongo.Collection, userID string) ([]Document, error)
	GetDocumentByID(ctx context.Context, coll *mongo.Collection, docID string) (*Document, error)
	DeleteDocumentByID(ctx context.Context, coll *mongo.Collection, docID string) error // Этот метод был пропущен
}

type RealStorage struct{}

func (s *RealStorage) InsertDocument(ctx context.Context, coll *mongo.Collection, doc Document) (primitive.ObjectID, error) {
	return InsertDocument(ctx, coll, doc)
}

func (s *RealStorage) GetDocumentsByUserID(ctx context.Context, coll *mongo.Collection, userID string) ([]Document, error) {
	return GetDocumentsByUserID(ctx, coll, userID)
}

func (s *RealStorage) GetDocumentByID(ctx context.Context, coll *mongo.Collection, docID string) (*Document, error) {
	return GetDocumentByID(ctx, coll, docID)
}

// Добавлен недостающий метод
func (s *RealStorage) DeleteDocumentByID(ctx context.Context, coll *mongo.Collection, docID string) error {
	return DeleteDocumentByID(ctx, coll, docID)
}

// Проверка реализации интерфейса
var _ Storage = (*RealStorage)(nil)
