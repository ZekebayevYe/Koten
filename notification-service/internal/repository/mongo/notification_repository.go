package mongo

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/erazz/outage-service/config"
	"github.com/erazz/outage-service/internal/domain"
)

type repository struct {
	col *mongo.Collection
}

func New(ctx context.Context, cfg config.Config) domain.NotificationRepository {
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	col := client.Database(cfg.MongoDB).Collection("notifications")
	idx := mongo.IndexModel{
		Keys: bson.D{{Key: "resource_type", Value: 1}, {Key: "location", Value: 1}, {Key: "start_time", Value: 1}},
	}
	_, _ = col.Indexes().CreateOne(ctx, idx)
	return &repository{col: col}
}

func (r *repository) Create(ctx context.Context, n *domain.Notification) (*domain.Notification, error) {
	n.CreatedAt = time.Now()
	res, err := r.col.InsertOne(ctx, n)
	if err != nil {
		return nil, err
	}
	n.ID = res.InsertedID.(primitive.ObjectID)
	return n, nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*domain.Notification, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var n domain.Notification
	err = r.col.FindOne(ctx, bson.M{"_id": oid}).Decode(&n)
	return &n, err
}

func (r *repository) List(ctx context.Context, filter domain.NotificationFilter) ([]domain.Notification, error) {
	query := bson.M{}

	if filter.ResourceType != "" {
		query["resource_type"] = filter.ResourceType
	}

	if filter.Location != "" {
		query["location"] = filter.Location
	}

	if !filter.DateFrom.IsZero() || !filter.DateTo.IsZero() {
		dateQuery := bson.M{}
		if !filter.DateFrom.IsZero() {
			dateQuery["$gte"] = filter.DateFrom
		}
		if !filter.DateTo.IsZero() {
			dateQuery["$lte"] = filter.DateTo
		}
		query["start_time"] = dateQuery
	}

	cur, err := r.col.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	var res []domain.Notification
	err = cur.All(ctx, &res)
	return res, err
}

func (r *repository) Update(ctx context.Context, n *domain.Notification) (*domain.Notification, error) {
	if n.ID.IsZero() {
		return nil, errors.New("missing id")
	}
	_, err := r.col.ReplaceOne(ctx, bson.M{"_id": n.ID}, n)
	return n, err
}

func (r *repository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.col.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}
