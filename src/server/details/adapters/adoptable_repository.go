package adapters

import (
	"encoding/json"
	"errors"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

const persistenceTypeAdoptable = "adoptables"

// AdoptableRepository ...
type AdoptableRepository struct {
	persistence Persistence
}

type persistedAdoptable struct {
	ID     string
	Name   string
	Labels []persistedLabel
}

type persistedLabel struct {
	Kind  string
	Value string
}

func createPersistedAdoptableFromAdoptable(entity *entities.Adoptable) *persistedAdoptable {
	return &persistedAdoptable{
		ID:     entity.ID.String(),
		Name:   entity.Name,
		Labels: createPersistedLabelListFromLabelList(entity.Labels),
	}
}

func createAdoptableFromPersistedAdoptable(pEntity *persistedAdoptable) *entities.Adoptable {
	return &entities.Adoptable{
		ID:     entities.BuildID(pEntity.ID),
		Name:   pEntity.Name,
		Labels: createLabelListFromPersistedLabelList(pEntity.Labels),
	}
}

func createPersistedAdoptableListFromAdoptableList(list []*entities.Adoptable) []*persistedAdoptable {
	output := make([]*persistedAdoptable, 0, 0)
	for _, item := range list {
		entity := createPersistedAdoptableFromAdoptable(item)
		output = append(output, entity)
	}

	return output
}

func createAdoptableListFromPersistedAdoptableList(list []*persistedAdoptable) []*entities.Adoptable {
	output := make([]*entities.Adoptable, 0, 0)
	for _, item := range list {
		entity := createAdoptableFromPersistedAdoptable(item)
		output = append(output, entity)
	}

	return output
}

func createPersistedLabelFromLabel(entity *entities.Label) persistedLabel {
	return persistedLabel{
		Kind:  entity.Kind,
		Value: entity.Value,
	}
}

func createPersistedLabelListFromLabelList(list []*entities.Label) []persistedLabel {
	output := make([]persistedLabel, 0, 0)
	for _, item := range list {
		entity := createPersistedLabelFromLabel(item)
		output = append(output, entity)
	}

	return output
}

func createLabelFromPersistedLabel(pEntity persistedLabel) *entities.Label {
	return &entities.Label{
		Kind:  pEntity.Kind,
		Value: pEntity.Value,
	}
}

func createLabelListFromPersistedLabelList(list []persistedLabel) []*entities.Label {
	output := make([]*entities.Label, 0, 0)
	for _, item := range list {
		entity := createLabelFromPersistedLabel(item)
		output = append(output, entity)
	}

	return output
}

func (t *persistedAdoptable) MarshalBinary() (data []byte, err error) {
	return json.Marshal(t)
}

func (t *persistedAdoptable) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &t); err != nil {
		glog.Error(err)
		return err
	}

	return nil
}

// Clone ...
func (t *persistedAdoptable) Clone() PersistedData {
	return &persistedAdoptable{}
}

// CreateAdoptableRepository ...
func CreateAdoptableRepository(persistence Persistence) *AdoptableRepository {
	return &AdoptableRepository{
		persistence: persistence,
	}
}

// GetAllAdoptables ...
func (r *AdoptableRepository) GetAllAdoptables() ([]*entities.Adoptable, error) {
	output := make([]*entities.Adoptable, 0, 0)
	proto := &persistedAdoptable{}
	items, _ := r.persistence.GetAll(persistenceTypeAdoptable, proto)
	for _, item := range items {
		entity := createAdoptableFromPersistedAdoptable(item.(*persistedAdoptable))
		output = append(output, entity)
	}

	return output, nil
}

// SaveAdoptable ...
func (r *AdoptableRepository) SaveAdoptable(entity *entities.Adoptable) (*entities.Adoptable, error) {
	pAdoptable := createPersistedAdoptableFromAdoptable(entity)
	if err := r.persistence.Create(persistenceTypeAdoptable, entity.ID.String(), pAdoptable); err != nil {
		glog.Error(err)
		return nil, err
	}

	return entity, nil
}

// FindAdoptableByID ...
func (r *AdoptableRepository) FindAdoptableByID(id string) (*entities.Adoptable, error) {
	proto := &persistedAdoptable{}
	pAdoptable, err := r.persistence.Find(persistenceTypeAdoptable, id, proto)
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("error getting adoptable")
	}
	if pAdoptable == nil {
		return nil, errors.New("adoptable doesn't exists")
	}

	entity := createAdoptableFromPersistedAdoptable(pAdoptable.(*persistedAdoptable))
	return entity, nil
}
