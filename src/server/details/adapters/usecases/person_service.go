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
	personRepository    personRepository
	adoptableRepository adoptableRepository
}

// CreatePersonService ...
func CreatePersonService(personRepository personRepository, adoptableRepository adoptableRepository) *PersonService {
	return &PersonService{
		personRepository:    personRepository,
		adoptableRepository: adoptableRepository,
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
func (s *PersonService) AddAdoptableToPerson(adoptableID string, personID string) (*entities.Person, error) {
	glog.Info(personID)
	person, _ := s.personRepository.FindPersonByID(personID)
	glog.Info(person)
	if person == nil {
		return nil, fmt.Errorf("person with id %s not found", personID)
	}
	adoptable, _ := s.adoptableRepository.FindAdoptableByID(adoptableID)
	person.AdoptAdoptable(adoptable)
	s.personRepository.SavePerson(person)

	return person, nil
}
