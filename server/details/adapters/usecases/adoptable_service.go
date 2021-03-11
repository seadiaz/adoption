package usecases

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/server/details/adapters/usecases/entities"
)

type adoptableRepository interface {
	GetAllAdoptables() ([]*entities.Adoptable, error)
	SaveAdoptable(entity *entities.Adoptable) (*entities.Adoptable, error)
	FindAdoptableByID(id string) (*entities.Adoptable, error)
}

// AdoptableService ...
type AdoptableService interface {
	GetAllAdoptables() ([]*entities.Adoptable, error)
	FindAdoptablesFilterByLabelKindAndValue(labelKind, labelValue string) ([]*entities.Adoptable, error)
	CreateAdoptable(name string) (*entities.Adoptable, error)
	FindAdoptable(id string) (*entities.Adoptable, error)
	AddLabelToAdoptable(labelKind string, labelValue string, adoptableID string) (*entities.Adoptable, error)
}

// AdoptableServiceExpert ...
type AdoptableServiceExpert struct {
	repository adoptableRepository
}

// CreateAdoptableService ...
func CreateAdoptableService(repository adoptableRepository) *AdoptableServiceExpert {
	return &AdoptableServiceExpert{
		repository: repository,
	}
}

// GetAllAdoptables ...
func (s *AdoptableServiceExpert) GetAllAdoptables() ([]*entities.Adoptable, error) {
	adoptables, _ := s.repository.GetAllAdoptables()
	return adoptables, nil
}

// FindAdoptablesFilterByLabelKindAndValue ...
func (s *AdoptableServiceExpert) FindAdoptablesFilterByLabelKindAndValue(labelKind, labelValue string) ([]*entities.Adoptable, error) {
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
func (s *AdoptableServiceExpert) CreateAdoptable(name string) (*entities.Adoptable, error) {
	adoptables, _ := s.repository.GetAllAdoptables()
	for _, item := range adoptables {
		if item.Name == name {
			return nil, fmt.Errorf("adoptable with name %s already exists", name)
		}
	}

	adoptable := entities.CreateAdoptableWithNameAndStrategy(name, entities.StrategyTypeSingle)
	adoptable, _ = s.repository.SaveAdoptable(adoptable)
	return adoptable, nil
}

// FindAdoptable ...
func (s *AdoptableServiceExpert) FindAdoptable(id string) (*entities.Adoptable, error) {
	adoptable, _ := s.repository.FindAdoptableByID(id)
	return adoptable, nil
}

// AddLabelToAdoptable ...
func (s *AdoptableServiceExpert) AddLabelToAdoptable(labelKind string, labelValue string, adoptableID string) (*entities.Adoptable, error) {
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
