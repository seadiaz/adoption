package adapters

import (
	"errors"

	"github.com/golang/glog"
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/src/details/adapters/usecases/entities"
)

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
	items := r.persistence.GetAll()
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
	team, err := r.persistence.Find(id)
	if err != nil {
		return nil, errors.New("error finding team by name")
	}

	return team.(*entities.Team), nil
}

// SaveTeam ...
func (r *TeamRepository) SaveTeam(entity *entities.Team) (*entities.Team, error) {
	team, _ := r.FindTeamByName(entity.Name)
	if team == nil {
		if err := r.persistence.Create(entity.ID, entity); err != nil {
			return nil, err
		}
	} else {
		if err := r.persistence.Update(entity.ID, entity); err != nil {
			return nil, err
		}
	}

	return entity, nil
}

// GetTeam ...
func (r *TeamRepository) GetTeam(id string) (*entities.Team, error) {
	res, err := r.persistence.Find(id)
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("error getting team")
	}
	if res == nil {
		return nil, errors.New("team doesn't exists")
	}

	return res.(*entities.Team), nil
}
