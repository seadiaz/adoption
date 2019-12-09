package usecases

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

type personRepository interface {
	FindPersonByID(id string) (*entities.Person, error)
	SavePerson(entity *entities.Person) (*entities.Person, error)
	GetAllPeople() ([]*entities.Person, error)
}

// PersonService ...
type PersonService struct {
	personRepository personRepository
	toolRepository   toolRepository
}

// CreatePersonService ...
func CreatePersonService(personRepository personRepository, toolRepository toolRepository) *PersonService {
	return &PersonService{
		personRepository: personRepository,
		toolRepository:   toolRepository,
	}
}

// GetAllPeople ...
func (s *PersonService) GetAllPeople() ([]*entities.Person, error) {
	people, _ := s.personRepository.GetAllPeople()
	return people, nil
}

// CreatePerson ...
func (s *PersonService) CreatePerson(name string, email string) (*entities.Person, error) {
	person := entities.CreatePersonWithNameAndEmail(name, email)
	_, err := s.personRepository.SavePerson(person)
	if err != nil {
		return nil, err
	}
	return person, nil
}

// AddAdoptableToPerson ...
func (s *PersonService) AddAdoptableToPerson(toolID string, personID string) (*entities.Person, error) {
	glog.Info(personID)
	person, _ := s.personRepository.FindPersonByID(personID)
	glog.Info(person)
	if person == nil {
		return nil, fmt.Errorf("person with id %s not found", personID)
	}
	tool, _ := s.toolRepository.FindAdoptableByID(toolID)
	person.AdoptAdoptable(tool)
	s.personRepository.SavePerson(person)

	return person, nil
}
