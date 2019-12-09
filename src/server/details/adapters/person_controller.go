package adapters

import (
	"net/http"

	"github.com/seadiaz/adoption/src/server/details/adapters/usecases"
)

var (
	createPersonRules = map[string]string{
		"name":  "required|string",
		"email": "required|email",
	}

	addAdoptableToPersonRules = map[string]string{
		"id": "required|uuid",
	}
)

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
	s.Router.HandleFunc("/people/{id}/tools", c.addAdoptableToPerson).Methods("POST")
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

func (c *PersonController) addAdoptableToPerson(w http.ResponseWriter, r *http.Request) {
	tool := &toolForm{}
	err := validateRequest(r, addAdoptableToPersonRules, tool)
	if err != nil {
		replyWithError(w, http.StatusBadRequest, err)
		return
	}

	id := getPathParam(r, "id")
	res, err := c.service.AddAdoptableToPerson(tool.ID, id)
	if err != nil {
		replyWithError(w, http.StatusBadRequest, err)
		return
	}

	output := CreatePersonResponseFromPerson(res)
	replyJSONResponse(w, output)
}
