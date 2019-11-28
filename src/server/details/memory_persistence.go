package details

import (
	"errors"
	"sync"

	"github.com/seadiaz/adoption/src/server/details/adapters"
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
func (p *MemoryPersistence) Create(id string, obj interface{}) error {
	if id == "" {
		return errors.New("you must provide an id")
	}
	p.memory.Store(id, obj)
	return nil
}

// Update ...
func (p *MemoryPersistence) Update(id string, obj interface{}) error {
	if id == "" {
		return errors.New("you must provide an id")
	}
	p.memory.Store(id, obj)
	return nil
}

// Delete ...
func (p *MemoryPersistence) Delete(id string) error {
	if id == "" {
		return errors.New("you must provide an id")
	}

	p.memory.Delete(id)
	return nil
}

// GetAll ...
func (p *MemoryPersistence) GetAll() []interface{} {
	list := make([]interface{}, 0, 0)
	p.memory.Range(func(_, value interface{}) bool {
		list = append(list, value)
		return true
	})

	return list
}

// Find ...
func (p *MemoryPersistence) Find(id string) (interface{}, error) {
	output, _ := p.memory.Load(id)
	return output, nil
}
