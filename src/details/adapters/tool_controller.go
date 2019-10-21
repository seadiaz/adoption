package adapters

import (
	"encoding/json"
	"net/http"

	usecases "github.com/seadiaz/adoption/src/details/adapters/use_cases"
)

// ToolController ...
type ToolController struct {
	service *usecases.ToolService
}

// CreateToolController ...
func CreateToolController(service *usecases.ToolService) ToolController {
	return ToolController{
		service: service,
	}
}

// AddRoutes ...
func (c *ToolController) AddRoutes(s Server) {
	s.Router.HandleFunc("/tools", c.getAllTools).Methods("GET")
	s.Router.HandleFunc("/tools", c.createTool).Methods("POST")
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
