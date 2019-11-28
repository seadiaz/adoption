package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/seadiaz/adoption/src/server/details/adapters/usecases"
)

var addMemberToTeamRules = map[string]string{
	"id": "required|uuid",
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
func (c *TeamController) AddRoutes(s Server) {
	s.Router.HandleFunc("/teams", c.getAllTeams).Methods("GET")
	s.Router.HandleFunc("/teams", c.createTeam).Methods("POST")
	s.Router.HandleFunc("/teams/{id}/people", c.addMemberToTeam).Methods("POST")
	s.Router.HandleFunc("/teams/{id}/people", c.getMembersFromTeam).Methods("GET")
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
	var entity map[string]interface{}
	json.NewDecoder(r.Body).Decode(&entity)
	name := entity["name"].(string)
	res, err := c.service.CreateTeam(name)
	if err != nil {
		replyWithError(w, http.StatusConflict, fmt.Errorf("error creating team. %s", err.Error()))
		return
	}
	output := CreateTeamResponseFromTeam(res)
	replyJSONResponse(w, output)
}

func (c *TeamController) addMemberToTeam(w http.ResponseWriter, r *http.Request) {
	person := &personForm{}
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
