package usecases

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/server/details/adapters/usecases/entities"
)

type teamRepository interface {
	GetAllTeams() ([]*entities.Team, error)
	SaveTeam(entity *entities.Team) (*entities.Team, error)
	FindTeamByName(name string) (*entities.Team, error)
	FindTeamByID(id string) (*entities.Team, error)
}

// TeamService ...
type TeamService struct {
	repository       teamRepository
	personRepository personRepository
}

// CreateTeamService ...
func CreateTeamService(repository teamRepository, personRepository personRepository) *TeamService {
	return &TeamService{
		repository:       repository,
		personRepository: personRepository,
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
	team, err := s.repository.FindTeamByID(id)
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
func (s *TeamService) AddMemberToTeam(personID string, teamID string) (*entities.Team, error) {
	team, _ := s.repository.FindTeamByID(teamID)
	person, _ := s.personRepository.FindPersonByID(personID)
	if memberExists(team, person) {
		glog.Warningf("person %s is already a member of %s", person.Name, team.Name)
		return team, nil
	}

	team.AddPerson(person)
	team, err := s.repository.SaveTeam(team)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return team, nil
}

func memberExists(team *entities.Team, person *entities.Person) bool {
	for _, item := range team.People {
		if item.ID.IsEqual(person.ID) {
			return true
		}
	}

	return false
}
