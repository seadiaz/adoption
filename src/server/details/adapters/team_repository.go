package adapters

import (
	"errors"

	"github.com/golang/glog"
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

const persistenceTypeTeam = "teams"

// TeamRepository ...
type TeamRepository struct {
	persistence Persistence
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
	items, _ := r.persistence.GetAll(persistenceTypeTeam)
	for _, item := range items {
		var entity *entities.Team
		mapstructure.Decode(item, &entity)
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

// FindTeam ...
func (r *TeamRepository) FindTeam(id string) (*entities.Team, error) {
	res, err := r.persistence.Find(persistenceTypeTeam, id)
	if err != nil {
		return nil, errors.New("error finding team by name")
	}
	if res == nil {
		return nil, errors.New("team doesn't exists")
	}

	return res.(*entities.Team), nil
}

// SaveTeam ...
func (r *TeamRepository) SaveTeam(entity *entities.Team) (*entities.Team, error) {
	team, _ := r.FindTeamByName(entity.Name)
	if team == nil {
		if err := r.persistence.Create(persistenceTypeTeam, entity.ID.String(), entity); err != nil {
			return nil, err
		}
	} else {
		if err := r.persistence.Update(persistenceTypeTeam, entity.ID.String(), entity); err != nil {
			return nil, err
		}
	}

	return entity, nil
}

// GetTeam ...
func (r *TeamRepository) GetTeam(id string) (*entities.Team, error) {
	res, err := r.persistence.Find(persistenceTypeTeam, id)
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("error getting team")
	}
	if res == nil {
		return nil, errors.New("team doesn't exists")
	}

	return res.(*entities.Team), nil
}
