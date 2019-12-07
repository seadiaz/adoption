package adapters

import (
	"encoding/json"
	"errors"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

const persistenceTypeTeam = "teams"

// TeamRepository ...
type TeamRepository struct {
	persistence Persistence
}

type persistedTeam struct {
	ID     string
	Name   string
	People []*persistedPerson
}

func createPersistedTeamFromTeam(entity *entities.Team) *persistedTeam {
	return &persistedTeam{
		ID:     entity.ID.String(),
		Name:   entity.Name,
		People: createPersistedPersonListFromPersonList(entity.People),
	}
}

func createTeamFromPersistedTeam(pEntity *persistedTeam) *entities.Team {
	return &entities.Team{
		ID:     entities.RecoverID(pEntity.ID),
		Name:   pEntity.Name,
		People: createPersonListFromPersistedPersonList(pEntity.People),
	}
}

func (t *persistedTeam) MarshalBinary() (data []byte, err error) {
	return json.Marshal(t)
}

func (t *persistedTeam) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &t); err != nil {
		glog.Error(err)
		return err
	}

	return nil
}

// CreateTeamRepository ...
func CreateTeamRepository(persistence Persistence) *TeamRepository {
	return &TeamRepository{
		persistence: persistence,
	}
}

// GetAllTeams ...
func (r *TeamRepository) GetAllTeams() ([]*entities.Team, error) {
	var output []*entities.Team
	proto := &persistedTeam{}
	items, _ := r.persistence.GetAll(persistenceTypeTeam, proto)
	for _, item := range items {
		entity := createTeamFromPersistedTeam(item.(*persistedTeam))
		output = append(output, entity)
	}

	return output, nil
}

// FindTeamByName ...
func (r *TeamRepository) FindTeamByName(name string) (*entities.Team, error) {
	items, err := r.GetAllTeams()
	if err != nil {
		return nil, errors.New("error finding team by name")
	}

	for _, item := range items {
		if item.Name == name {
			return item, nil
		}

	}

	return nil, nil
}

// FindTeamByID ...
func (r *TeamRepository) FindTeamByID(id string) (*entities.Team, error) {
	pTeam := &persistedTeam{}
	err := r.persistence.Find(persistenceTypeTeam, id, pTeam)
	if err != nil {
		return nil, errors.New("error finding team by name")
	}
	if pTeam == nil {
		return nil, errors.New("team doesn't exists")
	}

	entity := createTeamFromPersistedTeam(pTeam)
	return entity, nil
}

// SaveTeam ...
func (r *TeamRepository) SaveTeam(entity *entities.Team) (*entities.Team, error) {
	team, _ := r.FindTeamByName(entity.Name)
	pTeam := createPersistedTeamFromTeam(entity)
	if team == nil {
		if err := r.persistence.Create(persistenceTypeTeam, entity.ID.String(), pTeam); err != nil {
			return nil, err
		}
	} else {
		if err := r.persistence.Update(persistenceTypeTeam, entity.ID.String(), pTeam); err != nil {
			return nil, err
		}
	}

	return entity, nil
}
