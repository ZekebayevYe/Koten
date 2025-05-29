package cache

import (
	"auth-service/internal/domain"
	"github.com/patrickmn/go-cache"
	"time"
)

type UserCache struct {
	c *cache.Cache
}

func NewUserCache(defaultExpiration, cleanupInterval time.Duration) *UserCache {
	return &UserCache{
		c: cache.New(defaultExpiration, cleanupInterval),
	}
}

func (uc *UserCache) Get(email string) (*domain.User, bool) {
	val, found := uc.c.Get(email)
	if !found {
		return nil, false
	}
	user, ok := val.(*domain.User)
	return user, ok
}

func (uc *UserCache) Set(email string, user *domain.User) {
	uc.c.Set(email, user, cache.DefaultExpiration)
}

func (uc *UserCache) Invalidate(email string) {
	uc.c.Delete(email)
}
