package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/seadiaz/adoption/src/server/details/adapters/usecases"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

var (
	createToolRules = map[string]string{
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
	s.Router.HandleFunc("/tools/{id}", c.getToolByID).Methods("GET")
	s.Router.HandleFunc("/tools/{id}/adoption", c.calculateAdoptionForTool).Methods("GET")
	s.Router.HandleFunc("/tools/{id}/labels", c.addLabelToTool).Methods("POST")
}

func (c *ToolController) getAllTools(w http.ResponseWriter, r *http.Request) {
	kind, value := getQueryParamMapKeyValue(r, "labels")
	var res []*entities.Tool
	if kind != "" {
		res, _ = c.service.FindToolsFilterByLabelKindAndValue(kind, value)
	} else {
		res, _ = c.service.GetAllTools()
	}

	output := CreateToolResponseListFromToolList(res)
	replyJSONResponse(w, output)
}

func (c *ToolController) getToolByID(w http.ResponseWriter, r *http.Request) {
	id := getPathParam(r, "id")
	res, _ := c.service.FindTool(id)
	output := CreateToolResponseFromTool(res)
	replyJSONResponse(w, output)
}

func (c *ToolController) createTool(w http.ResponseWriter, r *http.Request) {
	tool := &toolForm{}
	err := validateRequest(r, createToolRules, tool)
	if err != nil {
		replyWithError(w, http.StatusBadRequest, err)
		return
	}

	res, err := c.service.CreateTool(tool.Name)
	if err != nil {
		replyWithError(w, http.StatusConflict, fmt.Errorf("error creating tool %s: %s", tool.Name, err.Error()))
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
	label := &labelForm{}
	err := validateRequest(r, labelRules, label)
	if err != nil {
		replyWithError(w, http.StatusBadRequest, err)
		return
	}

	var entity map[string]string
	json.NewDecoder(r.Body).Decode(&entity)
	id := getPathParam(r, "id")
	res, _ := c.service.AddLabelToTool(label.Kind, label.Value, id)
	output := CreateToolResponseFromTool(res)
	replyJSONResponse(w, output)
}
