package adapters

import (
	"errors"

	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/src/details/adapters/use_cases/entities"
)

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
	glog.Info("get all tools called")
	output := make([]*entities.Tool, 0, 0)
	items := r.persistence.GetAll()
	for _, item := range items {
		var entity *entities.Tool
		mapstructure.Decode(item, &entity)
		output = append(output, entity)
	}

	return output, nil
}

// SaveTool ...
func (r *ToolRepository) SaveTool(entity *entities.Tool) (*entities.Tool, error) {
	glog.Info("create tool called")
	if entity.ID == "" {
		entity.ID = uuid.New().String()
	}
	if err := r.persistence.Create(entity.ID, entity); err != nil {
		return nil, err
	}

	return entity, nil
}

// GetTool ...
func (r *ToolRepository) GetTool(id string) (*entities.Tool, error) {
	glog.Info("get tool called")
	res, err := r.persistence.Find(id)
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("error getting tool")
	}
	if res == nil {
		return nil, errors.New("tool doesn't exists")
	}

	return res.(*entities.Tool), nil
}
