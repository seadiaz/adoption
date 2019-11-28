package adapters

import (
	"encoding/json"
	"net/http"

	"github.com/seadiaz/adoption/src/server/details/adapters/usecases"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

var createPersonRules = map[string]string{
	"name":  "required|string",
	"email": "required|email",
}

type personForm struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// PersonController ...
type PersonController struct {
	service *usecases.PersonService
}

// CreatePersonController ...
func CreatePersonController(service *usecases.PersonService) PersonController {
	return PersonController{
		service: service,
	}
}

// AddRoutes ...
func (c *PersonController) AddRoutes(s Server) {
	s.Router.HandleFunc("/people", c.getAllPeople).Methods("GET")
	s.Router.HandleFunc("/people", c.createPerson).Methods("POST")
	s.Router.HandleFunc("/people/{id}/tools", c.addToolToPerson).Methods("POST")
}

func (c *PersonController) getAllPeople(w http.ResponseWriter, r *http.Request) {
	res, _ := c.service.GetAllPeople()
	output := CreatePersonResponseListFromPersonList(res)
	replyJSONResponse(w, output)
}

func (c *PersonController) createPerson(w http.ResponseWriter, r *http.Request) {
	person := &personForm{}
	err := validateRequest(r, createPersonRules, person)
	if err != nil {
		replyWithError(w, http.StatusBadRequest, err)
		return
	}
	res, _ := c.service.CreatePerson(person.Name, person.Email)
	output := CreatePersonResponseFromPerson(res)
	replyJSONResponse(w, output)
}

func (c *PersonController) addToolToPerson(w http.ResponseWriter, r *http.Request) {
	var entity *entities.Tool
	json.NewDecoder(r.Body).Decode(&entity)
	id := getPathParam(r, "id")
	res, _ := c.service.AddToolToPerson(entity, id)
	output := CreatePersonResponseFromPerson(res)
	replyJSONResponse(w, output)
}
