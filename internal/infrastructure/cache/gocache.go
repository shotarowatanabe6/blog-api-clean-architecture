package cache

import (
	"fmt"
	"time"

	"blog-api-clean-architecture/internal/domain/repository"

	gocache "github.com/patrickmn/go-cache"
)

type CacheRepository struct {
	Client *gocache.Cache
}

func NewCacheRepository() repository.IDBRepository {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	client := gocache.New(5*time.Minute, 10*time.Minute)
	return CacheRepository{client}
}

func (c CacheRepository) Get(key string) (*string, error) {
	value, ok := c.Client.Get(key)
	if !ok {
		return nil, nil
	}
	v, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("cache value for key %q has invalid type %T", key, value)
	}

	return &v, nil
}

func (c CacheRepository) Set(key string, value string) error {
	c.Client.Set(key, value, 5*time.Minute)
	return nil
}
