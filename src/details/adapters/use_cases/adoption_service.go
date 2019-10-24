package usecases

import (
	"errors"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/details/adapters/use_cases/entities"
)

// AdoptoinService ...
type AdoptoinService struct {
	toolRepository   toolRepository
	personRepository personRepository
}

// CreateAdoptionService ...
func CreateAdoptionService(toolRepository toolRepository, personRepository personRepository) *AdoptoinService {
	return &AdoptoinService{
		toolRepository:   toolRepository,
		personRepository: personRepository,
	}
}

// CalculateAdoptionForTool ...
func (s *AdoptoinService) CalculateAdoptionForTool(id string) (map[string]interface{}, error) {
	tool, err := s.toolRepository.GetTool(id)
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("calculate adoption for tool failed. " + err.Error())
	}
	people, err := s.personRepository.GetAllPeople()
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("calculate adoption for tool failed. " + err.Error())
	}

	adoption := entities.CreateAdoption()
	for _, person := range people {
		adoption.IncludePerson(person)
	}
	output := make(map[string]interface{})
	output["adoption"] = adoption.CalculateForTool(tool)

	return output, nil
}
