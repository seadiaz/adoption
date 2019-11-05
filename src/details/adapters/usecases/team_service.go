package usecases

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/details/adapters/usecases/entities"
)

type teamRepository interface {
	GetAllTeams() ([]*entities.Team, error)
	SaveTeam(entity *entities.Team) (*entities.Team, error)
	FindTeamByName(name string) (*entities.Team, error)
	FindTeam(id string) (*entities.Team, error)
	GetTeam(id string) (*entities.Team, error)
}

// TeamService ...
type TeamService struct {
	repository teamRepository
}

// CreateTeamService ...
func CreateTeamService(repository teamRepository) *TeamService {
	return &TeamService{
		repository: repository,
	}
}

// GetAllTeams ...
func (s *TeamService) GetAllTeams() ([]*entities.Team, error) {
	s.repository.GetAllTeams()
	teams, _ := s.repository.GetAllTeams()
	return teams, nil
}

// GetMembersFromTeam ...
func (s *TeamService) GetMembersFromTeam(id string) ([]*entities.Person, error) {
	team, err := s.repository.FindTeam(id)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return team.People, nil
}

// CreateTeam ...
func (s *TeamService) CreateTeam(name string) (*entities.Team, error) {
	teams, _ := s.repository.GetAllTeams()
	for _, item := range teams {
		if item.Name == name {
			return nil, fmt.Errorf("team with name %s already exists", name)
		}
	}

	team := entities.CreateTeamWithName(name)
	team, err := s.repository.SaveTeam(team)
	if err != nil {
		return nil, err
	}
	return team, nil
}

// AddMemberToTeam ...
func (s *TeamService) AddMemberToTeam(person *entities.Person, teamID string) (*entities.Team, error) {
	team, _ := s.repository.FindTeam(teamID)
	team.AddPerson(person)
	team, err := s.repository.SaveTeam(team)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return team, nil
}
