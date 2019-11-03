package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seadiaz/adoption/src/details/adapters/usecases"
)

// ToolController ...
type ToolController struct {
	service         *usecases.ToolService
	adoptionService *usecases.AdoptoinService
}

// CreateToolController ...
func CreateToolController(service *usecases.ToolService, adoptionService *usecases.AdoptoinService) ToolController {
	return ToolController{
		service:         service,
		adoptionService: adoptionService,
	}
}

// AddRoutes ...
func (c *ToolController) AddRoutes(s Server) {
	s.Router.HandleFunc("/tools", c.getAllTools).Methods("GET")
	s.Router.HandleFunc("/tools", c.createTool).Methods("POST")
	s.Router.HandleFunc("/tools/{id}/adoption", c.calculateAdoptionForTool).Methods("GET")
}

func (c *ToolController) getAllTools(w http.ResponseWriter, r *http.Request) {
	res, _ := c.service.GetAllTools()
	w.Header().Add("Content-Type", "application/json")
	output := CreateToolResponseListFromToolList(res)
	json.NewEncoder(w).Encode(output)
}

func (c *ToolController) createTool(w http.ResponseWriter, r *http.Request) {
	var entity map[string]interface{}
	json.NewDecoder(r.Body).Decode(&entity)
	name := entity["name"].(string)
	res, err := c.service.CreateTool(name)
	if err != nil {
		replyWithError(w, http.StatusConflict, fmt.Errorf("error creating tool. %s", err.Error()))
		return
	}
	output := CreateToolResponseFromTool(res)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (c *ToolController) calculateAdoptionForTool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	res, err := c.adoptionService.CalculateAdoptionForTool(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	output := CreateAdoptionResponseFromMap(res)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
