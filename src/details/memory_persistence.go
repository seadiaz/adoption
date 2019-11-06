package details

import (
	"errors"

	"github.com/seadiaz/adoption/src/details/adapters"
)

// MemoryPersistence is a memory implementantion of persistence
type MemoryPersistence struct {
	memory map[string]interface{}
}

// BuildMemoryPersistence ...
func BuildMemoryPersistence() adapters.Persistence {
	return &MemoryPersistence{
		memory: make(map[string]interface{}),
	}
}

// Create ...
func (p *MemoryPersistence) Create(id string, obj interface{}) error {
	if id == "" {
		return errors.New("you must provide an id")
	}
	p.memory[id] = obj
	return nil
}

// Update ...
func (p *MemoryPersistence) Update(id string, obj interface{}) error {
	if id == "" {
		return errors.New("you must provide an id")
	}
	p.memory[id] = obj
	return nil
}

// Delete ...
func (p *MemoryPersistence) Delete(id string) error {
	if id == "" {
		return errors.New("you must provide an id")
	}
	delete(p.memory, id)
	return nil
}

// GetAll ...
func (p *MemoryPersistence) GetAll() []interface{} {
	list := make([]interface{}, 0, 0)
	for _, value := range p.memory {
		list = append(list, value)
	}
	return list
}

// Find ...
func (p *MemoryPersistence) Find(id string) (interface{}, error) {
	output := p.memory[id]
	return output, nil
}
