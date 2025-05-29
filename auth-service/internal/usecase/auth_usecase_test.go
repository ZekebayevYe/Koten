package usecase_test

import (
	"context"
	"encoding/json"
	"testing"

	"auth-service/config"
	"auth-service/internal/domain"
	"auth-service/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateUser(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
func (m *MockRepo) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*domain.User), args.Error(1)
}
func (m *MockRepo) UpdateUser(ctx context.Context, email string, updated *domain.User) (*domain.User, error) {
	args := m.Called(ctx, email, updated)
	return args.Get(0).(*domain.User), args.Error(1)
}

type MockPublisher struct {
	mock.Mock
}

func (m *MockPublisher) Publish(ctx context.Context, subject string, data []byte) error {
	args := m.Called(ctx, subject, data)
	return args.Error(0)
}

func TestRegister_PublishesEvent(t *testing.T) {
	mockRepo := new(MockRepo)
	mockPublisher := new(MockPublisher)

	cfg := &config.Config{
		JWTSecret:    "testsecret",
		JWTExpiresIn: "60m",
	}

	user := &domain.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	mockRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	mockPublisher.On("Publish", mock.Anything, "user.registered", mock.MatchedBy(func(data []byte) bool {
		var evt map[string]string
		_ = json.Unmarshal(data, &evt)
		return evt["email"] == user.Email && evt["role"] == "user"
	})).Return(nil)

	uc := usecase.NewAuthUsecase(mockRepo, mockPublisher, cfg)

	token, err := uc.Register(context.Background(), user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	mockRepo.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
}
