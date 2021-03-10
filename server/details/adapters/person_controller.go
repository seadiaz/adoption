package adapters

import (
	"net/http"

	"github.com/seadiaz/adoption/server/details/adapters/usecases"
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

// PersonRequest ...
type PersonRequest struct {
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
func (c *PersonController) AddRoutes(r Router) {
	r.HandleFunc("/people", c.getAllPeople).Methods("GET")
	r.HandleFunc("/people", c.createPerson).Methods("POST")
	r.HandleFunc("/people/{id}/adoptables", c.addAdoptableToPerson).Methods("POST")
}

func (c *PersonController) getAllPeople(w http.ResponseWriter, r *http.Request) {
	res, _ := c.service.GetAllPeople()
	output := CreatePersonResponseListFromPersonList(res)
	replyJSONResponse(w, output)
}

func (c *PersonController) createPerson(w http.ResponseWriter, r *http.Request) {
	person := &PersonRequest{}
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
	adoptable := &adoptableForm{}
	err := validateRequest(r, addAdoptableToPersonRules, adoptable)
	if err != nil {
		replyWithError(w, http.StatusBadRequest, err)
		return
	}

	id := getPathParam(r, "id")
	res, err := c.service.AddAdoptableToPerson(adoptable.ID, id)
	if err != nil {
		replyWithError(w, http.StatusBadRequest, err)
		return
	}

	output := CreatePersonResponseFromPerson(res)
	replyJSONResponse(w, output)
}
