package adapters

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	usecases "github.com/seadiaz/adoption/src/details/adapters/use_cases"
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
	response, _ := c.service.GetAllTools()
	json.NewEncoder(w).Encode(response)
}

func (c *ToolController) createTool(w http.ResponseWriter, r *http.Request) {
	var entity map[string]interface{}
	json.NewDecoder(r.Body).Decode(&entity)
	name := entity["name"].(string)
	response, _ := c.service.CreateTool(name)
	json.NewEncoder(w).Encode(response)
}

func (c *ToolController) calculateAdoptionForTool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response, err := c.adoptionService.CalculateAdoptionForTool(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	json.NewEncoder(w).Encode(response)
}
