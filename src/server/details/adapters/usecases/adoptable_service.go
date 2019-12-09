package usecases

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

type toolRepository interface {
	GetAllAdoptables() ([]*entities.Adoptable, error)
	SaveAdoptable(entity *entities.Adoptable) (*entities.Adoptable, error)
	FindAdoptableByID(id string) (*entities.Adoptable, error)
}

// AdoptableService ...
type AdoptableService struct {
	repository toolRepository
}

// CreateAdoptableService ...
func CreateAdoptableService(repository toolRepository) *AdoptableService {
	return &AdoptableService{
		repository: repository,
	}
}

// GetAllAdoptables ...
func (s *AdoptableService) GetAllAdoptables() ([]*entities.Adoptable, error) {
	tools, _ := s.repository.GetAllAdoptables()
	return tools, nil
}

// FindAdoptablesFilterByLabelKindAndValue ...
func (s *AdoptableService) FindAdoptablesFilterByLabelKindAndValue(labelKind, labelValue string) ([]*entities.Adoptable, error) {
	tools, _ := s.repository.GetAllAdoptables()
	output := make([]*entities.Adoptable, 0, 0)
	for _, item := range tools {
		if item.HasLabelKindEqualToValue(labelKind, labelValue) {
			output = append(output, item)
		}
	}

	return output, nil
}

// CreateAdoptable ...
func (s *AdoptableService) CreateAdoptable(name string) (*entities.Adoptable, error) {
	tools, _ := s.repository.GetAllAdoptables()
	for _, item := range tools {
		if item.Name == name {
			return nil, fmt.Errorf("tool with name %s already exists", name)
		}
	}

	tool := entities.CreateAdoptableWithName(name)
	tool, _ = s.repository.SaveAdoptable(tool)
	return tool, nil
}

// FindAdoptable ...
func (s *AdoptableService) FindAdoptable(id string) (*entities.Adoptable, error) {
	tool, _ := s.repository.FindAdoptableByID(id)
	return tool, nil
}

// AddLabelToAdoptable ...
func (s *AdoptableService) AddLabelToAdoptable(labelKind string, labelValue string, toolID string) (*entities.Adoptable, error) {
	label := entities.CreateLabelWithKindAndValue(labelKind, labelValue)
	tool, _ := s.repository.FindAdoptableByID(toolID)
	tool.AddLabel(label)
	tool, err := s.repository.SaveAdoptable(tool)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return tool, nil
}
