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
func (p *MemoryPersistence) Create(kind string, id string, obj interface{}) error {
	if id == "" {
		return errors.New("you must provide an id")
	}
	p.memory.Store(id, obj)
	return nil
}

// Update ...
func (p *MemoryPersistence) Update(kind string, id string, obj interface{}) error {
	if id == "" {
		return errors.New("you must provide an id")
	}
	p.memory.Store(id, obj)
	return nil
}

// Delete ...
func (p *MemoryPersistence) Delete(kind string, id string) error {
	if id == "" {
		return errors.New("you must provide an id")
	}

	p.memory.Delete(id)
	return nil
}

// GetAll ...
func (p *MemoryPersistence) GetAll(kind string) ([]interface{}, error) {
	list := make([]interface{}, 0, 0)
	p.memory.Range(func(_, value interface{}) bool {
		list = append(list, value)
		return true
	})

	return list, nil
}

// Find ...
func (p *MemoryPersistence) Find(kind string, id string) (interface{}, error) {
	output, _ := p.memory.Load(id)
	return output, nil
}
