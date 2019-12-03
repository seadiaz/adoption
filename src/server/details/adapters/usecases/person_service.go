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

// AddToolToPerson ...
func (s *PersonService) AddToolToPerson(toolID string, personID string) (*entities.Person, error) {
	person, _ := s.personRepository.FindPerson(personID)
	tool, _ := s.toolRepository.FindTool(toolID)
	person.AdoptTool(tool)
	s.personRepository.SavePerson(person)

	return person, nil
}
