package details

import (
	"errors"

	"github.com/go-redis/redis/v7"
	"github.com/seadiaz/adoption/src/server/details/adapters"
)

// RedisPersistence is a redis implementantion of persistence
type RedisPersistence struct {
	client *redis.Client
}

// BuildRedisPersistence ...
func BuildRedisPersistence() adapters.Persistence {
	return &RedisPersistence{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

// Create ...
func (p *RedisPersistence) Create(id string, obj interface{}) error {
	return errors.New("not implemented")
}

// Update ...
func (p *RedisPersistence) Update(id string, obj interface{}) error {
	return errors.New("not implemented")
}

// Delete ...
func (p *RedisPersistence) Delete(id string) error {
	return errors.New("not implemented")
}

// GetAll ...
func (p *RedisPersistence) GetAll() []interface{} {
	return make([]interface{}, 0, 0)
}

// Find ...
func (p *RedisPersistence) Find(id string) (interface{}, error) {
	return nil, errors.New("not implemented")
}
