package usecases

import (
	"fmt"

	"github.com/seadiaz/adoption/src/details/adapters/usecases/entities"
)

type toolRepository interface {
	GetAllTools() ([]*entities.Tool, error)
	SaveTool(entity *entities.Tool) (*entities.Tool, error)
	GetTool(id string) (*entities.Tool, error)
}

// ToolService ...
type ToolService struct {
	repository toolRepository
}

// CreateToolService ...
func CreateToolService(repository toolRepository) *ToolService {
	return &ToolService{
		repository: repository,
	}
}

// GetAllTools ...
func (s *ToolService) GetAllTools() ([]*entities.Tool, error) {
	tools, _ := s.repository.GetAllTools()
	return tools, nil
}

// CreateTool ...
func (s *ToolService) CreateTool(name string) (*entities.Tool, error) {
	tools, _ := s.repository.GetAllTools()
	for _, item := range tools {
		if item.Name == name {
			return nil, fmt.Errorf("tool with name %s already exists", name)
		}
	}

	tool := entities.CreateToolWithName(name)
	tool, _ = s.repository.SaveTool(tool)
	return tool, nil
}
