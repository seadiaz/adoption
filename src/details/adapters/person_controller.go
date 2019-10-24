package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gookit/validate"
	"github.com/gorilla/mux"
	usecases "github.com/seadiaz/adoption/src/details/adapters/use_cases"
	"github.com/seadiaz/adoption/src/details/adapters/use_cases/entities"
)

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

type person struct {
	Name  string `json:"username"`
	Email string `json:"email"`
}

// AddRoutes ...
func (c *PersonController) AddRoutes(s Server) {
	s.Router.HandleFunc("/people", c.getAllPeople).Methods("GET")
	s.Router.HandleFunc("/people", c.createPerson).Methods("POST")
	s.Router.HandleFunc("/people/{id}/tools", c.addToolToPerson).Methods("POST")
}

func (c *PersonController) getAllPeople(w http.ResponseWriter, r *http.Request) {
	response, _ := c.service.GetAllPeople()
	json.NewEncoder(w).Encode(response)
}

func (c *PersonController) createPerson(w http.ResponseWriter, r *http.Request) {
	data, err := validate.FromRequest(r)
	if err != nil {
		panic(err)
	}
	v := data.Create()
	v.AddRule("name", "required")
	v.AddRule("email", "required")
	if v.Validate() {
		name, _ := data.Get("name")
		email, _ := data.Get("email")
		response, _ := c.service.CreatePerson(name.(string), email.(string))
		json.NewEncoder(w).Encode(response)
	} else {
		fmt.Println(v.Errors)
		fmt.Println(v.Errors.One())
		json.NewEncoder(w).Encode(v.Errors)
	}
}

func (c *PersonController) addToolToPerson(w http.ResponseWriter, r *http.Request) {
	var entity *entities.Tool
	json.NewDecoder(r.Body).Decode(&entity)
	vars := mux.Vars(r)
	id := vars["id"]
	person, _ := c.service.AddToolToPerson(entity, id)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}
