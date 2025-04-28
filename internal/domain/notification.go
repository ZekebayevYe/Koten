package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Title        string             `bson:"title"`
	Description  string             `bson:"description"`
	ResourceType string             `bson:"resource_type"`
	Location     string             `bson:"location"`
	StartTime    time.Time          `bson:"start_time"`
	EndTime      time.Time          `bson:"end_time"`
	CreatedAt    time.Time          `bson:"created_at"`
}

type NotificationFilter struct {
	ResourceType string
	Location     string
	DateFrom     time.Time
	DateTo       time.Time
}
