package adapters

import (
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

// AdoptionResponse ...
type AdoptionResponse struct {
	Adoptable        *AdoptableResponse `json:"adoptable"`
	Adoption         int                `json:"adoption"`
	Level            int                `json:"level"`
	Adopters         []*PersonResponse  `json:"adopters"`
	Absentees        []*PersonResponse  `json:"absentees"`
	TeamAdoption     int                `json:"team_adoption"`
	TeamOverallLevel int                `json:"team_overall_level"`
	TeamAdopters     []*AdoptionDetail  `json:"team_adopters"`
	TeamAbsentees    []*AdoptionDetail  `json:"team_absentees"`
}

// AdoptionDetail ...
type AdoptionDetail struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

// PersonResponse ...
type PersonResponse struct {
	ID         string               `json:"id"`
	Email      string               `json:"email"`
	Name       string               `json:"name"`
	Adoptables []*AdoptableResponse `json:"adoptables"`
}

// AdoptableResponse ...
type AdoptableResponse struct {
	ID     string           `json:"id"`
	Name   string           `json:"name,omitempty"`
	Labels []*LabelResponse `json:"labels,omitempty"`
}

// LabelResponse ...
type LabelResponse struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

// TeamResponse ...
type TeamResponse struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

// HealthResponse ...
type HealthResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

// CreatePersonResponseListFromPersonList ...
func CreatePersonResponseListFromPersonList(persons []*entities.Person) []*PersonResponse {
	output := make([]*PersonResponse, 0, 0)
	for _, item := range persons {
		output = append(output, CreatePersonResponseFromPerson(item))
	}

	return output
}

// CreatePersonResponseFromPerson ...
func CreatePersonResponseFromPerson(person *entities.Person) *PersonResponse {
	return &PersonResponse{
		ID:         person.ID.String(),
		Name:       person.Name,
		Email:      person.Email.String(),
		Adoptables: CreateAdoptableResponseListFromAdoptableList(person.Adoptables),
	}
}

// CreateAdoptableResponseListFromAdoptableList ...
func CreateAdoptableResponseListFromAdoptableList(adoptables []*entities.Adoptable) []*AdoptableResponse {
	output := make([]*AdoptableResponse, 0, 0)
	for _, item := range adoptables {
		output = append(output, CreateAdoptableResponseFromAdoptable(item))
	}

	return output
}

// CreateAdoptableResponseFromAdoptable ...
func CreateAdoptableResponseFromAdoptable(adoptable *entities.Adoptable) *AdoptableResponse {
	return &AdoptableResponse{
		ID:     adoptable.ID.String(),
		Name:   adoptable.Name,
		Labels: CreateLabelResponseListFromLabelList(adoptable.Labels),
	}
}

// CreateAdoptionResponseFromMap ...
func CreateAdoptionResponseFromMap(adoption map[string]interface{}) *AdoptionResponse {
	return &AdoptionResponse{
		Adoptable:        CreateAdoptableResponseFromAdoptable(adoption["adoptable"].(*entities.Adoptable)),
		Adoption:         adoption["level"].(int),
		Level:            adoption["level"].(int),
		Adopters:         CreatePersonResponseListFromPersonList(adoption["adopters"].([]*entities.Person)),
		Absentees:        CreatePersonResponseListFromPersonList(adoption["absentees"].([]*entities.Person)),
		TeamAdoption:     adoption["team_overall_level"].(int),
		TeamOverallLevel: adoption["team_overall_level"].(int),
		TeamAdopters:     CreateAdoptionDetailResponseListFromAdoptionTeamDetailList(adoption["team_adopters"].([]*entities.AdoptionTeamDetail)),
		TeamAbsentees:    CreateAdoptionDetailResponseListFromAdoptionTeamDetailList(adoption["team_absentees"].([]*entities.AdoptionTeamDetail)),
	}
}

// CreateTeamResponseFromTeam ...
func CreateTeamResponseFromTeam(team *entities.Team) *TeamResponse {
	return &TeamResponse{
		ID:   team.ID.String(),
		Name: team.Name,
	}
}

// CreateAdoptionDetailResponseFromAdoptionDetailTeam ...
func CreateAdoptionDetailResponseFromAdoptionDetailTeam(teamDetail *entities.AdoptionTeamDetail) *AdoptionDetail {
	return &AdoptionDetail{
		Name:  teamDetail.Team.Name,
		Level: teamDetail.Level,
	}
}

// CreateTeamResponseListFromTeamList ...
func CreateTeamResponseListFromTeamList(teams []*entities.Team) []*TeamResponse {
	output := make([]*TeamResponse, 0, 0)
	for _, item := range teams {
		output = append(output, CreateTeamResponseFromTeam(item))
	}

	return output
}

// CreateHealthResponseMap ...
func CreateHealthResponseMap(m map[string]string) *HealthResponse {
	output := &HealthResponse{
		Status:  m["status"],
		Message: m["message"],
	}

	return output
}

// CreateAdoptionDetailResponseListFromAdoptionTeamDetailList ...
func CreateAdoptionDetailResponseListFromAdoptionTeamDetailList(teamDetails []*entities.AdoptionTeamDetail) []*AdoptionDetail {
	output := make([]*AdoptionDetail, 0, 0)
	for _, item := range teamDetails {
		output = append(output, CreateAdoptionDetailResponseFromAdoptionDetailTeam(item))
	}

	return output
}

// CreateLabelResponseListFromLabelList ...
func CreateLabelResponseListFromLabelList(labels []*entities.Label) []*LabelResponse {
	output := make([]*LabelResponse, 0, 0)
	for _, item := range labels {
		output = append(output, CreateLabelResponseFromLabel(item))
	}

	return output
}

// CreateLabelResponseFromLabel ...
func CreateLabelResponseFromLabel(label *entities.Label) *LabelResponse {
	return &LabelResponse{
		Kind:  label.Kind,
		Value: label.Value,
	}
}
