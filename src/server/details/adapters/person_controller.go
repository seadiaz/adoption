package adapters

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
	"github.com/gookit/validate"
	"github.com/gorilla/mux"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
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
	res, _ := c.service.GetAllPeople()
	output := CreatePersonResponseListFromPersonList(res)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (c *PersonController) createPerson(w http.ResponseWriter, r *http.Request) {
	data, err := validateCreatePerson(r)
	if err != nil {
		replyWithError(w, http.StatusBadRequest, err)
	}
	res, _ := c.service.CreatePerson(data["name"].(string), data["email"].(string))
	output := CreatePersonResponseFromPerson(res)
	replyJSONResponse(w, output)
}

func validateCreatePerson(r *http.Request) (map[string]interface{}, error) {
	data, err := validate.FromRequest(r)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	v := data.Create()
	v.AddRule("name", "required")
	v.AddRule("email", "required")
	if !v.Validate() {
		return nil, v.Errors
	}
	output := make(map[string]interface{})
	output["name"], _ = data.Get("name")
	output["email"], _ = data.Get("email")
	return output, nil

}

func (c *PersonController) addToolToPerson(w http.ResponseWriter, r *http.Request) {
	var entity *entities.Tool
	json.NewDecoder(r.Body).Decode(&entity)
	vars := mux.Vars(r)
	id := vars["id"]
	res, _ := c.service.AddToolToPerson(entity, id)
	output := CreatePersonResponseFromPerson(res)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
