package adapters

import (
	"encoding/json"
	"errors"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/server/details/adapters/usecases/entities"
)

const persistenceTypePerson = "people"

// PersonRepository ...
type PersonRepository struct {
	persistence Persistence
}

type persistedPerson struct {
	ID         string
	Name       string
	Email      string
	Adoptables []*persistedAdoptable
}

func createPersistedPersonFromPerson(entity *entities.Person) *persistedPerson {
	return &persistedPerson{
		ID:         entity.ID.String(),
		Name:       entity.Name,
		Email:      entity.Email.String(),
		Adoptables: createPersistedAdoptableListFromAdoptableList(entity.Adoptables),
	}
}

func createPersistedPersonListFromPersonList(list []*entities.Person) []*persistedPerson {
	output := make([]*persistedPerson, 0, 0)
	for _, item := range list {
		entity := createPersistedPersonFromPerson(item)
		output = append(output, entity)
	}

	return output
}

func createPersonFromPersistedPerson(pEntity *persistedPerson) *entities.Person {
	return &entities.Person{
		ID:         entities.BuildID(pEntity.ID),
		Name:       pEntity.Name,
		Email:      entities.BuildEmail(pEntity.Email),
		Adoptables: createAdoptableListFromPersistedAdoptableList(pEntity.Adoptables),
	}
}

func createPersonListFromPersistedPersonList(pList []*persistedPerson) []*entities.Person {
	output := make([]*entities.Person, 0, 0)
	for _, item := range pList {
		entity := createPersonFromPersistedPerson(item)
		output = append(output, entity)
	}

	return output
}

func (p *persistedPerson) MarshalBinary() (data []byte, err error) {
	return json.Marshal(p)
}

func (p *persistedPerson) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &p); err != nil {
		glog.Error(err)
		return err
	}

	return nil
}

// Clone ...
func (p *persistedPerson) Clone() PersistedData {
	return &persistedPerson{}
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
	proto := &persistedPerson{}
	items, _ := r.persistence.GetAll(persistenceTypePerson, proto)
	for _, item := range items {
		entity := createPersonFromPersistedPerson(item.(*persistedPerson))
		output = append(output, entity)
	}

	return output, nil
}

// FindPersonByID ...
func (r *PersonRepository) FindPersonByID(id string) (*entities.Person, error) {
	proto := &persistedPerson{}
	pPerson, err := r.persistence.Find(persistenceTypePerson, id, proto)
	if err != nil {
		return nil, errors.New("error finding person by id")
	}
	if pPerson == nil {
		return nil, errors.New("person doesn't exists")
	}

	entity := createPersonFromPersistedPerson(pPerson.(*persistedPerson))
	return entity, nil
}

// SavePerson ...
func (r *PersonRepository) SavePerson(entity *entities.Person) (*entities.Person, error) {
	if !entity.Email.IsValid() {
		return nil, errors.New("person should have an email")
	}
	person, err := r.findPersonByEmail(entity.Email)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	if person == nil {
		pPerson := createPersistedPersonFromPerson(entity)
		if err := r.persistence.Create(persistenceTypePerson, entity.ID.String(), pPerson); err != nil {
			return nil, err
		}
	} else {
		entity.ID = person.ID
		pPerson := createPersistedPersonFromPerson(entity)
		if err := r.persistence.Update(persistenceTypePerson, entity.ID.String(), pPerson); err != nil {
			return nil, err
		}
	}

	return person, nil
}

func (r *PersonRepository) findPersonByEmail(email *entities.Email) (*entities.Person, error) {
	people, err := r.GetAllPeople()
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	for _, item := range people {
		if item.Email.IsEqual(email) {
			return item, nil
		}
	}

	return nil, nil
}
