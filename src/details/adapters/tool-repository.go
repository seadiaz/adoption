package adapters

import (
	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/src/details/adapters/use-cases/entities"
	"github.com/walmartdigital/katalog/src/server/persistence"
)

// ToolRepository ...
type ToolRepository struct {
	persistence persistence.Persistence
}

// CreateToolRepository ...
func CreateToolRepository(persistence persistence.Persistence) *ToolRepository {
	return &ToolRepository{
		persistence: persistence,
	}
}

// GetAllTools ...
func (r *ToolRepository) GetAllTools() []entities.Tool {
	glog.Info("get all tools called")
	var output []entities.Tool
	items := r.persistence.GetAll()
	for _, item := range items {
		var entity entities.Tool
		mapstructure.Decode(item, &entity)
		output = append(output, entity)
	}

	return output
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
