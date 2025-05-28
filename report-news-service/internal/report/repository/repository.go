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

// Insert создает новый отчет с правильным ObjectID и заполняет поле report_id
func (r *reportRepository) Insert(ctx context.Context, report *proto.Report) error {
	objID := primitive.NewObjectID()
	report.ReportId = objID.Hex() // Записываем строковое представление ObjectID

	doc := bson.M{
		"_id":         objID,
		"user_id":     report.UserId,
		"title":       report.Title,
		"description": report.Description,
		"created_at":  time.Now().Format(time.RFC3339),
		"updated_at":  time.Now().Format(time.RFC3339),
	}

	_, err := r.collection.InsertOne(ctx, doc)
	return err
}

// FindByUser возвращает все отчеты указанного пользователя
func (r *reportRepository) FindByUser(ctx context.Context, userID string) ([]*proto.Report, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reports []*proto.Report
	for cursor.Next(ctx) {
		var doc struct {
			ID          primitive.ObjectID `bson:"_id"`
			UserId      string             `bson:"user_id"`
			Title       string             `bson:"title"`
			Description string             `bson:"description"`
			CreatedAt   string             `bson:"created_at"`
			UpdatedAt   string             `bson:"updated_at"`
		}
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		report := &proto.Report{
			ReportId:    doc.ID.Hex(),
			UserId:      doc.UserId,
			Title:       doc.Title,
			Description: doc.Description,
			CreatedAt:   doc.CreatedAt,
			UpdatedAt:   doc.UpdatedAt,
		}
		reports = append(reports, report)
	}
	return reports, nil
}

// Update обновляет отчет по ObjectID и user_id
func (r *reportRepository) Update(ctx context.Context, req *proto.EditReportRequest) (*proto.Report, error) {
	objID, err := primitive.ObjectIDFromHex(req.ReportId)
	if err != nil {
		return nil, errors.New("invalid report_id format")
	}

	filter := bson.M{"_id": objID, "user_id": req.UserId}
	update := bson.M{
		"$set": bson.M{
			"title":       req.Title,
			"description": req.Description,
			"updated_at":  time.Now().Format(time.RFC3339),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := r.collection.FindOneAndUpdate(ctx, filter, update, opts)
	if result.Err() != nil {
		return nil, errors.New("report not found or update failed")
	}

	var doc struct {
		ID          primitive.ObjectID `bson:"_id"`
		UserId      string             `bson:"user_id"`
		Title       string             `bson:"title"`
		Description string             `bson:"description"`
		CreatedAt   string             `bson:"created_at"`
		UpdatedAt   string             `bson:"updated_at"`
	}

	if err := result.Decode(&doc); err != nil {
		return nil, err
	}

	updatedReport := &proto.Report{
		ReportId:    doc.ID.Hex(),
		UserId:      doc.UserId,
		Title:       doc.Title,
		Description: doc.Description,
		CreatedAt:   doc.CreatedAt,
		UpdatedAt:   doc.UpdatedAt,
	}

	return updatedReport, nil
}
