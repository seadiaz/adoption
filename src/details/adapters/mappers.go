package adapters

import (
	"github.com/seadiaz/adoption/src/details/adapters/usecases/entities"
)

// AdoptionResponse ...
type AdoptionResponse struct {
	Adoption  int               `json:"adoption"`
	Adopters  []*PersonResponse `json:"adopters"`
	Absentees []*PersonResponse `json:"absentees"`
}

// PersonResponse ...
type PersonResponse struct {
	Email string          `json:"email"`
	Name  string          `json:"name"`
	Tools []*ToolResponse `json:"tools"`
}

// ToolResponse ...
type ToolResponse struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
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
		Name:  person.Name,
		Email: person.Email,
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
		ID:   tool.ID,
		Name: tool.Name,
	}
}

// CreateAdoptionResponseFromMap ...
func CreateAdoptionResponseFromMap(adoption map[string]interface{}) *AdoptionResponse {
	return &AdoptionResponse{
		Adoption:  adoption["adoption"].(int),
		Adopters:  CreatePersonResponseListFromPersonList(adoption["adopters"].([]*entities.Person)),
		Absentees: CreatePersonResponseListFromPersonList(adoption["absentees"].([]*entities.Person)),
	}
}
