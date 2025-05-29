package handler

import (
	"context"
	"log"
	"time"

	"github.com/ZekebayevYe/documents-service/internal/broker"
	"github.com/ZekebayevYe/documents-service/internal/storage"
	documentpb "github.com/ZekebayevYe/documents-service/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocumentHandler struct {
	DB        *mongo.Collection
	Publisher broker.PublisherInterface // Используем интерфейс вместо конкретной реализации
	Storage   storage.Storage
	documentpb.UnimplementedDocumentServiceServer
}

func (h *DocumentHandler) Upload(ctx context.Context, req *documentpb.UploadRequest) (*documentpb.UploadResponse, error) {
	doc := storage.Document{
		UserID:    req.UserId,
		Filename:  req.Filename,
		Type:      req.Type,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}
	id, err := h.Storage.InsertDocument(ctx, h.DB, doc)
	if err != nil {
		return nil, err
	}

	// Публикация события о загрузке документа
	docID := id.Hex()
	if h.Publisher != nil {
		event := broker.DocumentEvent{
			EventType:  "uploaded",
			DocumentID: docID,
			UserID:     req.UserId,
			Filename:   req.Filename,
			Timestamp:  time.Now(),
		}

		if err := h.Publisher.PublishDocumentEvent(event); err != nil {
			log.Printf("Failed to publish event for document %s: %v", docID, err)
		}
	}

	return &documentpb.UploadResponse{
		DocumentId: docID,
		Status:     "uploaded",
	}, nil
}

func (h *DocumentHandler) GetUserDocuments(ctx context.Context, req *documentpb.UserRequest) (*documentpb.DocumentList, error) {
	docs, err := h.Storage.GetDocumentsByUserID(ctx, h.DB, req.UserId)
	if err != nil {
		return nil, err
	}
	resp := &documentpb.DocumentList{}
	for _, d := range docs {
		resp.Documents = append(resp.Documents, &documentpb.Document{
			Id:        d.ID.Hex(),
			Filename:  d.Filename,
			Type:      d.Type,
			CreatedAt: d.CreatedAt.Format(time.RFC3339),
		})
	}
	return resp, nil
}

func (h *DocumentHandler) Download(ctx context.Context, req *documentpb.DocumentID) (*documentpb.DocumentContent, error) {
	doc, err := h.Storage.GetDocumentByID(ctx, h.DB, req.DocumentId)
	if err != nil {
		return nil, err
	}
	return &documentpb.DocumentContent{
		Filename: doc.Filename,
		Content:  doc.Content,
	}, nil
}

func (h *DocumentHandler) Delete(ctx context.Context, req *documentpb.DocumentID) (*documentpb.DeleteStatus, error) {
	err := h.Storage.DeleteDocumentByID(ctx, h.DB, req.DocumentId)
	if err != nil {
		return nil, err
	}
	return &documentpb.DeleteStatus{Status: "deleted"}, nil
}
