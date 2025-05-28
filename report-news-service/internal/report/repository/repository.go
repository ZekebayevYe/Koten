package repository

import (
	"context"
	"errors"
	"time"

	"reportnewsservice/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReportRepository interface {
	Insert(ctx context.Context, report *proto.Report) error
	FindByUser(ctx context.Context, userID string) ([]*proto.Report, error)
	Update(ctx context.Context, req *proto.EditReportRequest) (*proto.Report, error)
}

type reportRepository struct {
	collection *mongo.Collection
}

func NewReportRepository(db *mongo.Database) ReportRepository {
	return &reportRepository{
		collection: db.Collection("reports"),
	}
}

func (r *reportRepository) Insert(ctx context.Context, report *proto.Report) error {
	_, err := r.collection.InsertOne(ctx, report)
	return err
}

func (r *reportRepository) FindByUser(ctx context.Context, userID string) ([]*proto.Report, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reports []*proto.Report
	for cursor.Next(ctx) {
		var report proto.Report
		if err := cursor.Decode(&report); err != nil {
			return nil, err
		}
		reports = append(reports, &report)
	}
	return reports, nil
}

func (r *reportRepository) Update(ctx context.Context, req *proto.EditReportRequest) (*proto.Report, error) {
	objID, err := primitive.ObjectIDFromHex(req.ReportId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID, "user_id": req.UserId}
	update := bson.M{
		"$set": bson.M{
			"title":       req.Title,
			"description": req.Description,
			"updated_at":  time.Now().Format(time.RFC3339),
		},
	}
	result := r.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		return nil, errors.New("report not found or update failed")
	}
	var updated proto.Report
	if err := result.Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}
