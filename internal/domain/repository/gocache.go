package repository

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

type CacheRepository struct {
	Client *gocache.Cache
}

func NewCacheRepository() IDBRepository {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	client := gocache.New(5*time.Minute, 10*time.Minute)
	return CacheRepository{client}
}

type IDBRepository interface {
	Get(key string) (*string, error)
	Set(key string, value string) error
}

func (c CacheRepository) Get(key string) (*string, error) {
	value, ok := c.Client.Get(key)
	if !ok {
		return nil, nil
	}
	v := value.(string)

	return &v, nil
}

func (c CacheRepository) Set(key string, value string) error {
	c.Client.Set(key, value, 5*time.Minute)
	return nil
}
