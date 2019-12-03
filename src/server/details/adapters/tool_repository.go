package adapters

import (
	"errors"
	"fmt"

	"github.com/golang/glog"
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

const persistenceTypeTool = "tools"

// ToolRepository ...
type ToolRepository struct {
	persistence Persistence
}

// CreateToolRepository ...
func CreateToolRepository(persistence Persistence) *ToolRepository {
	return &ToolRepository{
		persistence: persistence,
	}
}

// GetAllTools ...
func (r *ToolRepository) GetAllTools() ([]*entities.Tool, error) {
	output := make([]*entities.Tool, 0, 0)
	items, _ := r.persistence.GetAll(persistenceTypeTool)
	for _, item := range items {
		var entity *entities.Tool
		mapstructure.Decode(item, &entity)
		output = append(output, entity)
	}

	return output, nil
}

// SaveTool ...
func (r *ToolRepository) SaveTool(entity *entities.Tool) (*entities.Tool, error) {
	if entity.ID.String() == "" {
		return nil, fmt.Errorf("ID is missing for tool %s", entity.Name)
	}
	if err := r.persistence.Create(persistenceTypeTool, entity.ID.String(), entity); err != nil {
		return nil, err
	}

	return entity, nil
}

// GetTool ...
func (r *ToolRepository) GetTool(id string) (*entities.Tool, error) {
	res, err := r.persistence.Find(persistenceTypeTool, id)
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("error getting tool")
	}
	if res == nil {
		return nil, errors.New("tool doesn't exists")
	}

	return res.(*entities.Tool), nil
}

// FindTool ...
func (r *ToolRepository) FindTool(id string) (*entities.Tool, error) {
	res, err := r.persistence.Find(persistenceTypeTool, id)
	if err != nil {
		return nil, errors.New("error finding tool")
	}
	if res == nil {
		return nil, errors.New("tool doesn't exists")
	}

	return res.(*entities.Tool), nil
}
