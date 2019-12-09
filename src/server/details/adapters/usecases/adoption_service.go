package usecases

import (
	"errors"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

// AdoptoinService ...
type AdoptoinService struct {
	adoptableRepository adoptableRepository
	personRepository    personRepository
	teamRepository      teamRepository
}

// CreateAdoptionService ...
func CreateAdoptionService(adoptableRepository adoptableRepository, personRepository personRepository, teamRepository teamRepository) *AdoptoinService {
	return &AdoptoinService{
		adoptableRepository: adoptableRepository,
		personRepository:    personRepository,
		teamRepository:      teamRepository,
	}
}

// CalculateAdoptionForAdoptable ...
func (s *AdoptoinService) CalculateAdoptionForAdoptable(id string) (map[string]interface{}, error) {
	adoptable, err := s.adoptableRepository.FindAdoptableByID(id)
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("calculate adoption for adoptable failed. " + err.Error())
	}

	people, err := s.personRepository.GetAllPeople()
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("calculate adoption for adoptable failed. " + err.Error())
	}

	teams, err := s.teamRepository.GetAllTeams()
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("calculate adoption for adoptable failed. " + err.Error())
	}

	adoption := entities.CreateAdoption()
	for _, item := range people {
		adoption.IncludePerson(item)
	}
	for _, item := range teams {
		adoption.IncludeTeam(item)
	}
	output := make(map[string]interface{})
	output["adoption"] = adoption.CalculateForAdoptable(adoptable)
	output["adopters"] = adoption.FilterAdoptersForAdoptable(adoptable)
	output["absentees"] = adoption.FilterAbsenteesForAdoptable(adoptable)
	output["team_adoption"] = adoption.CalculateTeamForAdoptable(adoptable)
	output["team_adopters"] = adoption.FilterTeamAdoptersForAdoptable(adoptable)
	output["team_absentees"] = adoption.FilterTeamAbsenteesForAdoptable(adoptable)

	return output, nil
}
