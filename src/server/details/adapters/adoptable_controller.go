package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/seadiaz/adoption/src/server/details/adapters/usecases"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

var (
	createAdoptableRules = map[string]string{
		"name": "required|string",
	}

	labelRules = map[string]string{
		"kind":  "required|string",
		"value": "required|string",
	}
)

type toolForm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type labelForm struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

// AdoptableController ...
type AdoptableController struct {
	service         *usecases.AdoptableService
	adoptionService *usecases.AdoptoinService
}

// CreateAdoptableController ...
func CreateAdoptableController(service *usecases.AdoptableService, adoptionService *usecases.AdoptoinService) AdoptableController {
	return AdoptableController{
		service:         service,
		adoptionService: adoptionService,
	}
}

// AddRoutes ...
func (c *AdoptableController) AddRoutes(s Server) {
	s.Router.HandleFunc("/tools", c.getAllAdoptables).Methods("GET")
	s.Router.HandleFunc("/tools", c.createAdoptable).Methods("POST")
	s.Router.HandleFunc("/tools/{id}", c.getAdoptableByID).Methods("GET")
	s.Router.HandleFunc("/tools/{id}/adoption", c.calculateAdoptionForAdoptable).Methods("GET")
	s.Router.HandleFunc("/tools/{id}/labels", c.addLabelToAdoptable).Methods("POST")
}

func (c *AdoptableController) getAllAdoptables(w http.ResponseWriter, r *http.Request) {
	kind, value := getQueryParamMapKeyValue(r, "labels")
	var res []*entities.Adoptable
	if kind != "" {
		res, _ = c.service.FindAdoptablesFilterByLabelKindAndValue(kind, value)
	} else {
		res, _ = c.service.GetAllAdoptables()
	}

	output := CreateAdoptableResponseListFromAdoptableList(res)
	replyJSONResponse(w, output)
}

func (c *AdoptableController) getAdoptableByID(w http.ResponseWriter, r *http.Request) {
	id := getPathParam(r, "id")
	res, _ := c.service.FindAdoptable(id)
	output := CreateAdoptableResponseFromAdoptable(res)
	replyJSONResponse(w, output)
}

func (c *AdoptableController) createAdoptable(w http.ResponseWriter, r *http.Request) {
	tool := &toolForm{}
	err := validateRequest(r, createAdoptableRules, tool)
	if err != nil {
		replyWithError(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.service.CreateAdoptable(tool.Name)
	if err != nil {
		replyWithError(w, http.StatusConflict, fmt.Errorf("error creating tool %s: %s", tool.Name, err.Error()))
		return
	}
	output := CreateAdoptableResponseFromAdoptable(res)
	replyJSONResponse(w, output)
}

func (c *AdoptableController) calculateAdoptionForAdoptable(w http.ResponseWriter, r *http.Request) {
	id := getPathParam(r, "id")
	res, err := c.adoptionService.CalculateAdoptionForAdoptable(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	output := CreateAdoptionResponseFromMap(res)
	replyJSONResponse(w, output)
}

func (c *AdoptableController) addLabelToAdoptable(w http.ResponseWriter, r *http.Request) {
	label := &labelForm{}
	err := validateRequest(r, labelRules, label)
	if err != nil {
		replyWithError(w, http.StatusBadRequest, err)
		return
	}

	var entity map[string]string
	json.NewDecoder(r.Body).Decode(&entity)
	id := getPathParam(r, "id")
	res, _ := c.service.AddLabelToAdoptable(label.Kind, label.Value, id)
	output := CreateAdoptableResponseFromAdoptable(res)
	replyJSONResponse(w, output)
}
