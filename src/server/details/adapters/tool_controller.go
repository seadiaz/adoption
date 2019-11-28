package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/seadiaz/adoption/src/server/details/adapters/usecases"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
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
	s.Router.HandleFunc("/tools/{id}/labels", c.addLabelToTool).Methods("POST")
}

func (c *ToolController) getAllTools(w http.ResponseWriter, r *http.Request) {
	res, _ := c.service.GetAllTools()
	output := CreateToolResponseListFromToolList(res)
	replyJSONResponse(w, output)
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
	replyJSONResponse(w, output)
}

func (c *ToolController) calculateAdoptionForTool(w http.ResponseWriter, r *http.Request) {
	id := getPathParam(r, "id")
	res, err := c.adoptionService.CalculateAdoptionForTool(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	output := CreateAdoptionResponseFromMap(res)
	replyJSONResponse(w, output)
}

func (c *ToolController) addLabelToTool(w http.ResponseWriter, r *http.Request) {
	var entity map[string]string
	json.NewDecoder(r.Body).Decode(&entity)
	id := getPathParam(r, "id")
	label := entities.CreateLabelWithKindAndValue(entity["kind"], entity["value"])
	res, _ := c.service.AddLabelToTool(label, id)
	output := CreateToolResponseFromTool(res)
	replyJSONResponse(w, output)
}
