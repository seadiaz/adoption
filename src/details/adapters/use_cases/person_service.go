package usecases

import (
	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/details/adapters/use_cases/entities"
)

type personRepository interface {
	FindPerson(id string) (*entities.Person, error)
	SavePerson(entity *entities.Person) error
	GetAllPeople() []entities.Person
}

// PersonService ...
type PersonService struct {
	repository personRepository
}

// CreatePersonService ...
func CreatePersonService(repository personRepository) *PersonService {
	return &PersonService{
		repository: repository,
	}
}

// GetAllPeople ...
func (s *PersonService) GetAllPeople() ([]entities.Person, error) {
	people := s.repository.GetAllPeople()
	return people, nil
}

// CreatePerson ...
func (s *PersonService) CreatePerson(name string, email string) (*entities.Person, error) {
	glog.Info("create person called")
	person := entities.CreatePersonWithNameAndEmail(name, email)
	err := s.repository.SavePerson(person)
	if err != nil {
		return nil, err
	}
	return person, nil
}

// AddToolToPerson ...
func (s *PersonService) AddToolToPerson(tool *entities.Tool, personId string) (*entities.Person, error) {
	person, _ := s.repository.FindPerson(personId)
	person.AdoptTool(tool)
	s.repository.SavePerson(person)

	return person, nil
}
