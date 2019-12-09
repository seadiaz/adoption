package usecases

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

type adoptableRepository interface {
	GetAllAdoptables() ([]*entities.Adoptable, error)
	SaveAdoptable(entity *entities.Adoptable) (*entities.Adoptable, error)
	FindAdoptableByID(id string) (*entities.Adoptable, error)
}

// AdoptableService ...
type AdoptableService struct {
	repository adoptableRepository
}

// CreateAdoptableService ...
func CreateAdoptableService(repository adoptableRepository) *AdoptableService {
	return &AdoptableService{
		repository: repository,
	}
}

// GetAllAdoptables ...
func (s *AdoptableService) GetAllAdoptables() ([]*entities.Adoptable, error) {
	adoptables, _ := s.repository.GetAllAdoptables()
	return adoptables, nil
}

// FindAdoptablesFilterByLabelKindAndValue ...
func (s *AdoptableService) FindAdoptablesFilterByLabelKindAndValue(labelKind, labelValue string) ([]*entities.Adoptable, error) {
	adoptables, _ := s.repository.GetAllAdoptables()
	output := make([]*entities.Adoptable, 0, 0)
	for _, item := range adoptables {
		if item.HasLabelKindEqualToValue(labelKind, labelValue) {
			output = append(output, item)
		}
	}

	return output, nil
}

// CreateAdoptable ...
func (s *AdoptableService) CreateAdoptable(name string) (*entities.Adoptable, error) {
	adoptables, _ := s.repository.GetAllAdoptables()
	for _, item := range adoptables {
		if item.Name == name {
			return nil, fmt.Errorf("adoptable with name %s already exists", name)
		}
	}

	adoptable := entities.CreateAdoptableWithName(name)
	adoptable, _ = s.repository.SaveAdoptable(adoptable)
	return adoptable, nil
}

// FindAdoptable ...
func (s *AdoptableService) FindAdoptable(id string) (*entities.Adoptable, error) {
	adoptable, _ := s.repository.FindAdoptableByID(id)
	return adoptable, nil
}

// AddLabelToAdoptable ...
func (s *AdoptableService) AddLabelToAdoptable(labelKind string, labelValue string, adoptableID string) (*entities.Adoptable, error) {
	label := entities.CreateLabelWithKindAndValue(labelKind, labelValue)
	adoptable, _ := s.repository.FindAdoptableByID(adoptableID)
	adoptable.AddLabel(label)
	adoptable, err := s.repository.SaveAdoptable(adoptable)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return adoptable, nil
}
