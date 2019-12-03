package adapters

import (
	"errors"

	"github.com/golang/glog"
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

const persistenceTypePerson = "people"

// PersonRepository ...
type PersonRepository struct {
	persistence Persistence
}

// CreatePersonRepository ...
func CreatePersonRepository(persistence Persistence) *PersonRepository {
	return &PersonRepository{
		persistence: persistence,
	}
}

// GetAllPeople ...
func (r *PersonRepository) GetAllPeople() ([]*entities.Person, error) {
	var output []*entities.Person
	items, _ := r.persistence.GetAll(persistenceTypePerson)
	for _, item := range items {
		var entity *entities.Person
		mapstructure.Decode(item, &entity)
		output = append(output, entity)
	}

	return output, nil
}

// FindPerson ...
func (r *PersonRepository) FindPerson(id string) (*entities.Person, error) {
	var output entities.Person
	item, _ := r.persistence.Find(persistenceTypePerson, id)
	mapstructure.Decode(item, &output)

	return &output, nil
}

// SavePerson ...
func (r *PersonRepository) SavePerson(entity *entities.Person) (*entities.Person, error) {
	if entity.Email == "" {
		return nil, errors.New("person should have an email")
	}
	person, err := r.findPersonByEmail(entity.Email)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	if person == nil {
		if err := r.persistence.Create(persistenceTypePerson, entity.ID.String(), entity); err != nil {
			return nil, err
		}
	} else {
		entity.ID = person.ID
		if err := r.persistence.Update(persistenceTypePerson, person.ID.String(), entity); err != nil {
			return nil, err
		}
	}

	return person, nil
}

func (r *PersonRepository) findPersonByEmail(email string) (*entities.Person, error) {
	people, err := r.GetAllPeople()
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	for _, item := range people {
		if item.Email == email {
			return item, nil
		}
	}

	return nil, nil
}
