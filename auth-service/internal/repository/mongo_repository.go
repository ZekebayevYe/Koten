package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"auth-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database, collectionName string) *MongoUserRepository {
	return &MongoUserRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoUserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	count, err := r.collection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		fmt.Println("[MongoRepo] CountDocuments error:", err)
		return err
	}
	if count > 0 {
		return errors.New("user with this email already exists")
	}

	fmt.Println("[MongoRepo] CreateUser email:", user.Email, "password:", user.Password)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = r.collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("[MongoRepo] InsertOne error:", err)
		return err
	}
	fmt.Println("[MongoRepo] CreateUser success for email:", user.Email)
	return nil
}

func (r *MongoUserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		fmt.Println("[MongoRepo] FindOne error for email:", email, "error:", err)
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	fmt.Println("[MongoRepo] GetUserByEmail fetched: email =", user.Email, "password =", user.Password)
	return &user, nil
}

func (r *MongoUserRepository) UpdateUser(ctx context.Context, email string, updateData *domain.User) (*domain.User, error) {
	update := bson.M{
		"$set": bson.M{
			"full_name":  updateData.FullName,
			"house":      updateData.House,
			"street":     updateData.Street,
			"apartment":  updateData.Apartment,
			"updated_at": time.Now(),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedUser domain.User
	err := r.collection.FindOneAndUpdate(ctx, bson.M{"email": email}, update, opts).Decode(&updatedUser)
	if err != nil {
		fmt.Println("[MongoRepo] FindOneAndUpdate error for email:", email, "error:", err)
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	fmt.Println("[MongoRepo] UpdateUser updated: email =", updatedUser.Email)
	return &updatedUser, nil
}