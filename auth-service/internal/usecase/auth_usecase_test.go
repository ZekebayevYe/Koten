package usecase_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"
 
	"auth-service/config"
	"auth-service/internal/domain"
	"auth-service/internal/usecase"
	"auth-service/internal/usecase/cache"

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

	cache := cache.NewUserCache(5*time.Minute, 10*time.Minute) 

	mockRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	mockPublisher.On("Publish", mock.Anything, "user.registered", mock.MatchedBy(func(data []byte) bool {
		var evt map[string]string
		_ = json.Unmarshal(data, &evt)
		return evt["email"] == user.Email && evt["role"] == "user"
	})).Return(nil)

	uc := usecase.NewAuthUsecase(mockRepo, mockPublisher, cfg, cache)

	token, err := uc.Register(context.Background(), user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	mockRepo.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
}

func TestGetProfile_UsesCache(t *testing.T) {
	mockRepo := new(MockRepo)
	mockPublisher := new(MockPublisher)

	cfg := &config.Config{}
	cache := cache.NewUserCache(5*time.Minute, 10*time.Minute) 
	uc := usecase.NewAuthUsecase(mockRepo, mockPublisher, cfg, cache) 

	email := "cached@example.com"
	expectedUser := &domain.User{
		Email:    email,
		FullName: "Test User",
	}

	mockRepo.On("GetUserByEmail", mock.Anything, email).Return(expectedUser, nil).Once()

	user1, err1 := uc.GetProfile(context.Background(), email)
	assert.NoError(t, err1)
	assert.Equal(t, expectedUser, user1)

	user2, err2 := uc.GetProfile(context.Background(), email)
	assert.NoError(t, err2)
	assert.Equal(t, expectedUser, user2)

	mockRepo.AssertExpectations(t)
}
