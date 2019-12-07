package details

import (
	"github.com/go-redis/redis/v7"
	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/server/details/adapters"
)

// RedisPersistence is a redis implementantion of persistence
type RedisPersistence struct {
	client *redis.Client
}

// BuildRedisPersistence ...
func BuildRedisPersistence() adapters.Persistence {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ping(client)
	return &RedisPersistence{
		client: client,
	}
}

func ping(client *redis.Client) {
	_, err := client.Ping().Result()
	if err != nil {
		glog.Fatal(err)
	}

	glog.Info("connected to redis")
}

// Create ...
func (p *RedisPersistence) Create(kind string, id string, obj adapters.PersistedData) error {
	if _, err := p.client.HSet(kind, id, obj).Result(); err != nil {
		glog.Error(err)
		return err
	}

	return nil
}

// Update ...
func (p *RedisPersistence) Update(kind string, id string, obj adapters.PersistedData) error {
	if _, err := p.client.HSet(kind, id, obj).Result(); err != nil {
		glog.Error(err)
		return err
	}

	return nil
}

// Delete ...
func (p *RedisPersistence) Delete(kind string, id string) error {
	if _, err := p.client.HDel(kind, id).Result(); err != nil {
		glog.Error(err)
		return err
	}

	return nil
}

// GetAll ...
func (p *RedisPersistence) GetAll(kind string, proto adapters.PersistedData) ([]interface{}, error) {
	res, err := p.client.HGetAll(kind).Result()
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	list := make([]interface{}, 0, 0)
	for _, item := range res {
		entity := proto.Clone()
		entity.UnmarshalBinary([]byte(item))
		glog.Info(entity)
		list = append(list, entity)
	}
	return list, nil
}

// Find ...
func (p *RedisPersistence) Find(kind string, id string, proto adapters.PersistedData) (interface{}, error) {
	res, err := p.client.HGet(kind, id).Result()
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	entity := proto.Clone()
	entity.UnmarshalBinary([]byte(res))
	return entity, nil
}
