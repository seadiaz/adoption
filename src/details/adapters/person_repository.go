package adapters

import (
	"errors"

	"github.com/golang/glog"
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/src/details/adapters/use_cases/entities"
)

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
func (r *PersonRepository) GetAllPeople() ([]entities.Person, error) {
	glog.Info("get all tools called")
	var output []entities.Person
	items := r.persistence.GetAll()
	for _, item := range items {
		var entity entities.Person
		mapstructure.Decode(item, &entity)
		output = append(output, entity)
	}

	return output, nil
}

// FindPerson ...
func (r *PersonRepository) FindPerson(id string) (*entities.Person, error) {
	glog.Info("get person called")
	var output entities.Person
	item, _ := r.persistence.Find(id)
	mapstructure.Decode(item, &output)

	return &output, nil
}

// SavePerson ...
func (r *PersonRepository) SavePerson(entity *entities.Person) error {
	glog.Info("create person called")
	if entity.Email == "" {
		return errors.New("person should have an email")
	}
	person, _ := r.FindPerson(entity.Email)
	if person == nil {
		if err := r.persistence.Create(entity.Email, entity); err != nil {
			return err
		}
	} else {
		if err := r.persistence.Update(entity.Email, entity); err != nil {
			return err
		}
	}

	return nil
}
