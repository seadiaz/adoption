package adapters

import (
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

// AdoptionResponse ...
type AdoptionResponse struct {
	Adoption      int               `json:"adoption"`
	Adopters      []*PersonResponse `json:"adopters"`
	Absentees     []*PersonResponse `json:"absentees"`
	TeamAdoption  int               `json:"team_adoption"`
	TeamAdopters  []*TeamResponse   `json:"team_adopters"`
	TeamAbsentees []*TeamResponse   `json:"team_absentees"`
}

// PersonResponse ...
type PersonResponse struct {
	ID    string          `json:"id"`
	Email string          `json:"email"`
	Name  string          `json:"name"`
	Tools []*ToolResponse `json:"tools"`
}

// ToolResponse ...
type ToolResponse struct {
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
		ID:    person.ID.String(),
		Name:  person.Name,
		Email: person.Email.String(),
		Tools: CreateToolResponseListFromToolList(person.Tools),
	}
}

// CreateToolResponseListFromToolList ...
func CreateToolResponseListFromToolList(tools []*entities.Tool) []*ToolResponse {
	output := make([]*ToolResponse, 0, 0)
	for _, item := range tools {
		output = append(output, CreateToolResponseFromTool(item))
	}

	return output
}

// CreateToolResponseFromTool ...
func CreateToolResponseFromTool(tool *entities.Tool) *ToolResponse {
	return &ToolResponse{
		ID:     tool.ID.String(),
		Name:   tool.Name,
		Labels: CreateLabelResponseListFromLabelList(tool.Labels),
	}
}

// CreateAdoptionResponseFromMap ...
func CreateAdoptionResponseFromMap(adoption map[string]interface{}) *AdoptionResponse {
	return &AdoptionResponse{
		Adoption:      adoption["adoption"].(int),
		Adopters:      CreatePersonResponseListFromPersonList(adoption["adopters"].([]*entities.Person)),
		Absentees:     CreatePersonResponseListFromPersonList(adoption["absentees"].([]*entities.Person)),
		TeamAdoption:  adoption["team_adoption"].(int),
		TeamAdopters:  CreateTeamResponseListFromTeamList(adoption["team_adopters"].([]*entities.Team)),
		TeamAbsentees: CreateTeamResponseListFromTeamList(adoption["team_absentees"].([]*entities.Team)),
	}
}

// CreateTeamResponseFromTeam ...
func CreateTeamResponseFromTeam(team *entities.Team) *TeamResponse {
	return &TeamResponse{
		ID:   team.ID.String(),
		Name: team.Name,
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
