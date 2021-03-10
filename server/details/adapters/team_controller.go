package adapters

import (
	"fmt"
	"net/http"

	"github.com/seadiaz/adoption/server/details/adapters/usecases"
)

var (
	addMemberToTeamRules = map[string]string{
		"id": "required|uuid",
	}

	createTeamRules = map[string]string{
		"name": "required|string",
	}
)

type teamForm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// TeamController ...
type TeamController struct {
	service *usecases.TeamService
}

// CreateTeamController ...
func CreateTeamController(service *usecases.TeamService) *TeamController {
	return &TeamController{
		service: service,
	}
}

// AddRoutes ...
func (c *TeamController) AddRoutes(r Router) {
	r.HandleFunc("/teams", c.getAllTeams).Methods("GET")
	r.HandleFunc("/teams", c.createTeam).Methods("POST")
	r.HandleFunc("/teams/{id}/people", c.addMemberToTeam).Methods("POST")
	r.HandleFunc("/teams/{id}/people", c.getMembersFromTeam).Methods("GET")
}

func (c *TeamController) getAllTeams(w http.ResponseWriter, r *http.Request) {
	res, _ := c.service.GetAllTeams()
	output := CreateTeamResponseListFromTeamList(res)
	replyJSONResponse(w, output)
}

func (c *TeamController) getMembersFromTeam(w http.ResponseWriter, r *http.Request) {
	id := getPathParam(r, "id")
	res, _ := c.service.GetMembersFromTeam(id)
	output := CreatePersonResponseListFromPersonList(res)
	replyJSONResponse(w, output)
}

func (c *TeamController) createTeam(w http.ResponseWriter, r *http.Request) {
	team := &teamForm{}
	err := validateRequest(r, createTeamRules, team)
	if err != nil {
		replyWithError(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.service.CreateTeam(team.Name)
	if err != nil {
		replyWithError(w, http.StatusConflict, fmt.Errorf("error creating team. %s", err.Error()))
		return
	}
	output := CreateTeamResponseFromTeam(res)
	replyJSONResponse(w, output)
}

func (c *TeamController) addMemberToTeam(w http.ResponseWriter, r *http.Request) {
	person := &PersonRequest{}
	err := validateRequest(r, addMemberToTeamRules, person)
	if err != nil {
		replyWithError(w, http.StatusBadRequest, err)
		return
	}
	id := getPathParam(r, "id")
	res, _ := c.service.AddMemberToTeam(person.ID, id)
	output := CreateTeamResponseFromTeam(res)
	replyJSONResponse(w, output)
}
