package usecases

import "github.com/seadiaz/adoption/src/details/adapters/use-cases/entities"

type toolRepository interface {
	GetAllTools() []entities.Tool
	SaveTool(entity *entities.Tool) (*entities.Tool, error)
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
func (s *ToolService) GetAllTools() ([]entities.Tool, error) {
	tools := s.repository.GetAllTools()
	return tools, nil
}

// CreateTool ...
func (s *ToolService) CreateTool(name string) (*entities.Tool, error) {
	tool := entities.CreateToolWithName(name)
	tool, _ = s.repository.SaveTool(tool)
	return tool, nil
}
