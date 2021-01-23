package details

import (
	"errors"
	"strings"
	"sync"

	"github.com/seadiaz/adoption/server/details/adapters"
)

// MemoryPersistence is a memory implementantion of persistence
type MemoryPersistence struct {
	memory *sync.Map
}

// BuildMemoryPersistence ...
func BuildMemoryPersistence() adapters.Persistence {
	return &MemoryPersistence{
		memory: &sync.Map{},
	}
}

// Create ...
func (p *MemoryPersistence) Create(kind string, id string, obj adapters.PersistedData) error {
	if id == "" {
		return errors.New("you must provide an id")
	}
	res, _ := obj.MarshalBinary()
	p.memory.Store(kind+"-"+id, string(res))
	return nil
}

// Update ...
func (p *MemoryPersistence) Update(kind string, id string, obj adapters.PersistedData) error {
	if id == "" {
		return errors.New("you must provide an id")
	}
	res, _ := obj.MarshalBinary()
	p.memory.Store(kind+"-"+id, string(res))
	return nil
}

// Delete ...
func (p *MemoryPersistence) Delete(kind string, id string) error {
	if id == "" {
		return errors.New("you must provide an id")
	}

	p.memory.Delete(kind + "-" + id)
	return nil
}

// GetAll ...
func (p *MemoryPersistence) GetAll(kind string, proto adapters.PersistedData) ([]interface{}, error) {
	list := make([]interface{}, 0, 0)
	p.memory.Range(func(key, value interface{}) bool {
		if strings.Contains(key.(string), kind) {
			item := proto.Clone()
			item.UnmarshalBinary([]byte(value.(string)))
			list = append(list, item)
		}
		return true
	})

	return list, nil
}

// Find ...
func (p *MemoryPersistence) Find(kind string, id string, proto adapters.PersistedData) (interface{}, error) {
	res, _ := p.memory.Load(kind + "-" + id)
	if res == nil {
		return nil, nil
	}
	entity := proto.Clone()
	entity.UnmarshalBinary([]byte(res.(string)))
	return entity, nil
}

// MemoryTransaction ...
type MemoryTransaction struct{}

// BeginTransaction ...
func (p *MemoryPersistence) BeginTransaction() adapters.Transaction {
	return &MemoryTransaction{}
}

// Commit ...
func (t *MemoryTransaction) Commit() error {
	return nil
}

// Rollback ...
func (t *MemoryTransaction) Rollback() error {
	return nil
}
