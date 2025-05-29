package handler

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ZekebayevYe/documents-service/internal/broker"
	"github.com/ZekebayevYe/documents-service/internal/storage"
	documentpb "github.com/ZekebayevYe/documents-service/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MockStorage - мок хранилища
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) InsertDocument(ctx context.Context, coll *mongo.Collection, doc storage.Document) (primitive.ObjectID, error) {
	args := m.Called(ctx, coll, doc)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}

func (m *MockStorage) GetDocumentsByUserID(ctx context.Context, coll *mongo.Collection, userID string) ([]storage.Document, error) {
	args := m.Called(ctx, coll, userID)
	return args.Get(0).([]storage.Document), args.Error(1)
}

func (m *MockStorage) GetDocumentByID(ctx context.Context, coll *mongo.Collection, docID string) (*storage.Document, error) {
	args := m.Called(ctx, coll, docID)
	// Возвращаем указатель на Document
	return args.Get(0).(*storage.Document), args.Error(1)
}

func (m *MockStorage) DeleteDocumentByID(ctx context.Context, coll *mongo.Collection, docID string) error {
	args := m.Called(ctx, coll, docID)
	return args.Error(0)
}

// MockPublisher - мок издателя
type MockPublisher struct {
	mock.Mock
}

func (m *MockPublisher) PublishDocumentEvent(event broker.DocumentEvent) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockPublisher) Close() {}

func TestDocumentHandler_Upload(t *testing.T) {
	// Создаем моки
	mockStorage := new(MockStorage)
	mockPublisher := new(MockPublisher)

	// Настраиваем ожидания
	expectedID := primitive.NewObjectID()
	mockStorage.On("InsertDocument", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedID, nil)
	mockPublisher.On("PublishDocumentEvent", mock.Anything).Return(nil)

	// Создаем обработчик с моками
	handler := &DocumentHandler{
		DB:        nil,
		Publisher: mockPublisher,
		Storage:   mockStorage,
	}

	// Тестовые данные
	req := &documentpb.UploadRequest{
		UserId:   "test-user",
		Filename: "test.txt",
		Type:     "text/plain",
		Content:  []byte("test content"),
	}

	// Вызываем метод
	resp, err := handler.Upload(context.Background(), req)

	// Проверяем результат
	assert.NoError(t, err)
	assert.Equal(t, expectedID.Hex(), resp.DocumentId)
	assert.Equal(t, "uploaded", resp.Status)

	// Проверяем вызовы
	mockStorage.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
}

func TestDocumentHandler_GetUserDocuments(t *testing.T) {
	mockStorage := new(MockStorage)
	now := time.Now()

	// Настраиваем ожидания
	expectedDocs := []storage.Document{
		{
			ID:        primitive.NewObjectID(),
			UserID:    "test-user",
			Filename:  "file1.txt",
			Type:      "text/plain",
			CreatedAt: now,
		},
	}
	mockStorage.On("GetDocumentsByUserID", mock.Anything, mock.Anything, "test-user").
		Return(expectedDocs, nil)

	handler := &DocumentHandler{
		DB:      nil,
		Storage: mockStorage,
	}

	// Тестовые данные
	req := &documentpb.UserRequest{
		UserId: "test-user",
	}

	// Вызываем метод
	resp, err := handler.GetUserDocuments(context.Background(), req)

	// Проверяем результат
	assert.NoError(t, err)
	assert.Len(t, resp.Documents, 1)
	assert.Equal(t, expectedDocs[0].ID.Hex(), resp.Documents[0].Id)
	assert.Equal(t, expectedDocs[0].Filename, resp.Documents[0].Filename)
	assert.Equal(t, expectedDocs[0].Type, resp.Documents[0].Type)
	assert.Equal(t, expectedDocs[0].CreatedAt.Format(time.RFC3339), resp.Documents[0].CreatedAt)

	// Проверяем вызовы
	mockStorage.AssertExpectations(t)
}

func TestDocumentHandler_Download(t *testing.T) {
	mockStorage := new(MockStorage)
	docID := "60d5ec9ef3b7a12588e0f2a1"
	expectedDoc := &storage.Document{ // Теперь указатель
		Filename: "test.txt",
		Content:  []byte("test content"),
	}

	// Настраиваем ожидания
	mockStorage.On("GetDocumentByID", mock.Anything, mock.Anything, docID).
		Return(expectedDoc, nil)

	handler := &DocumentHandler{
		DB:      nil,
		Storage: mockStorage,
	}

	// Тестовые данные
	req := &documentpb.DocumentID{
		DocumentId: docID,
	}

	// Вызываем метод
	resp, err := handler.Download(context.Background(), req)

	// Проверяем результат
	assert.NoError(t, err)
	assert.Equal(t, expectedDoc.Filename, resp.Filename)
	assert.Equal(t, expectedDoc.Content, resp.Content)

	// Проверяем вызовы
	mockStorage.AssertExpectations(t)
}

func TestDocumentHandler_Delete(t *testing.T) {
	mockStorage := new(MockStorage)
	docID := "60d5ec9ef3b7a12588e0f2a1"

	// Настраиваем ожидания
	mockStorage.On("DeleteDocumentByID", mock.Anything, mock.Anything, docID).
		Return(nil)

	handler := &DocumentHandler{
		DB:      nil,
		Storage: mockStorage,
	}

	// Тестовые данные
	req := &documentpb.DocumentID{
		DocumentId: docID,
	}

	// Вызываем метод
	resp, err := handler.Delete(context.Background(), req)

	// Проверяем результат
	assert.NoError(t, err)
	assert.Equal(t, "deleted", resp.Status)

	// Проверяем вызовы
	mockStorage.AssertExpectations(t)
}

func TestDocumentHandler_Upload_PublishError(t *testing.T) {
	mockStorage := new(MockStorage)
	mockPublisher := new(MockPublisher)

	expectedID := primitive.NewObjectID()
	mockStorage.On("InsertDocument", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedID, nil)
	mockPublisher.On("PublishDocumentEvent", mock.Anything).Return(fmt.Errorf("publish error"))

	handler := &DocumentHandler{
		DB:        nil,
		Publisher: mockPublisher,
		Storage:   mockStorage,
	}

	req := &documentpb.UploadRequest{
		UserId:   "test-user",
		Filename: "test.txt",
		Type:     "text/plain",
		Content:  []byte("test content"),
	}

	resp, err := handler.Upload(context.Background(), req)

	assert.NoError(t, err) // Ошибка публикации не должна влиять на основной результат
	assert.Equal(t, expectedID.Hex(), resp.DocumentId)
	assert.Equal(t, "uploaded", resp.Status)

	mockStorage.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
}

func TestDocumentHandler_GetUserDocuments_Error(t *testing.T) {
	mockStorage := new(MockStorage)

	mockStorage.On("GetDocumentsByUserID", mock.Anything, mock.Anything, "test-user").
		Return([]storage.Document{}, fmt.Errorf("database error"))

	handler := &DocumentHandler{
		DB:      nil,
		Storage: mockStorage,
	}

	req := &documentpb.UserRequest{
		UserId: "test-user",
	}

	_, err := handler.GetUserDocuments(context.Background(), req)

	assert.Error(t, err)
	mockStorage.AssertExpectations(t)
}
