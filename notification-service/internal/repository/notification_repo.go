package repository

import (
	"context"
	"github.com/ZekebayevYe/notification-service/internal/app"
	"github.com/ZekebayevYe/notification-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Subscriber struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Email  string             `bson:"email"`
	Street string             `bson:"street"`
	House  string             `bson:"house"`
}

type NotificationEntity struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title"`
	Message string             `bson:"message"`
	SendAt  int64              `bson:"send_at"`
	Street  string             `bson:"street"`
}

type NotificationRepo struct {
	db *mongo.Database
}

func NewNotificationRepo(client *mongo.Client) *NotificationRepo {
	return &NotificationRepo{db: client.Database("notifications")}
}

func (r *NotificationRepo) AddSubscriber(ctx context.Context, s model.Subscriber) error {
	coll := r.db.Collection("subscribers")
	_, err := coll.InsertOne(ctx, s)
	if mongo.IsDuplicateKeyError(err) {
		return app.ErrAlreadySubscribed
	}
	return err
}

func (r *NotificationRepo) RemoveSubscriber(ctx context.Context, email string) error {
	coll := r.db.Collection("subscribers")
	_, err := coll.DeleteOne(ctx, bson.M{"email": email})
	return err
}

func (r *NotificationRepo) ListSubscribers(ctx context.Context) ([]string, error) {
	coll := r.db.Collection("subscribers")
	cur, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var subs []string
	for cur.Next(ctx) {
		var s Subscriber
		_ = cur.Decode(&s)
		subs = append(subs, s.Email)
	}
	return subs, nil
}

func (r *NotificationRepo) SaveNotification(ctx context.Context, n model.Notification) error {
	coll := r.db.Collection("notifications")
	_, err := coll.InsertOne(ctx, NotificationEntity{
		Title:   n.Title,
		Message: n.Message,
		SendAt:  n.SendAt,
	})
	return err
}
func (r *NotificationRepo) ListSubscribersByStreet(ctx context.Context, street string) ([]string, error) {
	coll := r.db.Collection("subscribers")
	filter := bson.M{"street": street}

	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var subs []string
	for cur.Next(ctx) {
		var s Subscriber
		if err := cur.Decode(&s); err != nil {
			continue
		}
		subs = append(subs, s.Email)
	}
	return subs, nil
}
func (r *NotificationRepo) GetAllNotifications(ctx context.Context) ([]model.Notification, error) {
	coll := r.db.Collection("notifications")
	cur, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var result []model.Notification
	for cur.Next(ctx) {
		var e NotificationEntity
		_ = cur.Decode(&e)
		result = append(result, model.Notification{
			Title:   e.Title,
			Message: e.Message,
			SendAt:  e.SendAt,
			Street:  e.Street,
		})
	}
	return result, nil
}
