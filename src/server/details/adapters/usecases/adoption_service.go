package usecases

import (
	"errors"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

// AdoptoinService ...
type AdoptoinService struct {
	toolRepository   toolRepository
	personRepository personRepository
	teamRepository   teamRepository
}

// CreateAdoptionService ...
func CreateAdoptionService(toolRepository toolRepository, personRepository personRepository, teamRepository teamRepository) *AdoptoinService {
	return &AdoptoinService{
		toolRepository:   toolRepository,
		personRepository: personRepository,
		teamRepository:   teamRepository,
	}
}

// CalculateAdoptionForTool ...
func (s *AdoptoinService) CalculateAdoptionForTool(id string) (map[string]interface{}, error) {
	tool, err := s.toolRepository.FindToolByID(id)
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("calculate adoption for tool failed. " + err.Error())
	}

	people, err := s.personRepository.GetAllPeople()
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("calculate adoption for tool failed. " + err.Error())
	}

	teams, err := s.teamRepository.GetAllTeams()
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("calculate adoption for tool failed. " + err.Error())
	}

	adoption := entities.CreateAdoption()
	for _, item := range people {
		adoption.IncludePerson(item)
	}
	for _, item := range teams {
		adoption.IncludeTeam(item)
	}
	output := make(map[string]interface{})
	output["adoption"] = adoption.CalculateForTool(tool)
	output["adopters"] = adoption.FilterAdoptersForTool(tool)
	output["absentees"] = adoption.FilterAbsenteesForTool(tool)
	output["team_adoption"] = adoption.CalculateTeamForTool(tool)
	output["team_adopters"] = adoption.FilterTeamAdoptersForTool(tool)
	output["team_absentees"] = adoption.FilterTeamAbsenteesForTool(tool)

	return output, nil
}
