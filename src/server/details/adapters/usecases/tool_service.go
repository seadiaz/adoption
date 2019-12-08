package usecases

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

type toolRepository interface {
	GetAllTools() ([]*entities.Tool, error)
	SaveTool(entity *entities.Tool) (*entities.Tool, error)
	FindToolByID(id string) (*entities.Tool, error)
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

// FindToolsFilterByLabelKindAndValue ...
func (s *ToolService) FindToolsFilterByLabelKindAndValue(labelKind, labelValue string) ([]*entities.Tool, error) {
	tools, _ := s.repository.GetAllTools()
	output := make([]*entities.Tool, 0, 0)
	for _, item := range tools {
		if item.HasLabelKindEqualToValue(labelKind, labelValue) {
			output = append(output, item)
		}
	}

	return output, nil
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

// FindTool ...
func (s *ToolService) FindTool(id string) (*entities.Tool, error) {
	tool, _ := s.repository.FindToolByID(id)
	return tool, nil
}

// AddLabelToTool ...
func (s *ToolService) AddLabelToTool(labelKind string, labelValue string, toolID string) (*entities.Tool, error) {
	label := entities.CreateLabelWithKindAndValue(labelKind, labelValue)
	tool, _ := s.repository.FindToolByID(toolID)
	tool.AddLabel(label)
	tool, err := s.repository.SaveTool(tool)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return tool, nil
}
