package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seadiaz/adoption/src/details/adapters/usecases"
	"github.com/seadiaz/adoption/src/details/adapters/usecases/entities"
)

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
	w.Header().Add("Content-Type", "application/json")
	output := CreateTeamResponseListFromTeamList(res)
	json.NewEncoder(w).Encode(output)
}

func (c *TeamController) getMembersFromTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	res, _ := c.service.GetMembersFromTeam(id)
	w.Header().Add("Content-Type", "application/json")
	output := CreatePersonResponseListFromPersonList(res)
	json.NewEncoder(w).Encode(output)
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
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (c *TeamController) addMemberToTeam(w http.ResponseWriter, r *http.Request) {
	var entity *entities.Person
	json.NewDecoder(r.Body).Decode(&entity)
	vars := mux.Vars(r)
	id := vars["id"]
	res, _ := c.service.AddMemberToTeam(entity, id)
	output := CreateTeamResponseFromTeam(res)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
