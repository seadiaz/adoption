package usecases

import (
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

type personRepository interface {
	FindPerson(id string) (*entities.Person, error)
	SavePerson(entity *entities.Person) (*entities.Person, error)
	GetAllPeople() ([]*entities.Person, error)
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
func (s *PersonService) GetAllPeople() ([]*entities.Person, error) {
	people, _ := s.repository.GetAllPeople()
	return people, nil
}

// CreatePerson ...
func (s *PersonService) CreatePerson(name string, email string) (*entities.Person, error) {
	person := entities.CreatePersonWithNameAndEmail(name, email)
	_, err := s.repository.SavePerson(person)
	if err != nil {
		return nil, err
	}
	return person, nil
}

// AddToolToPerson ...
func (s *PersonService) AddToolToPerson(tool *entities.Tool, personID string) (*entities.Person, error) {
	person, _ := s.repository.FindPerson(personID)
	person.AdoptTool(tool)
	s.repository.SavePerson(person)

	return person, nil
}
