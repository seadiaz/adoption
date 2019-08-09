package adapters

import (
	"encoding/json"
	"net/http"

	"github.com/seadiaz/adoption/src/details/adapters/use-cases/entities"
)

// ToolController ...
type ToolController struct {
	repository *ToolRepository
}

// CreateToolController ...
func CreateToolController(repository *ToolRepository) ToolController {
	return ToolController{
		repository: repository,
	}
}

// AddRoutes ...
func (c *ToolController) AddRoutes(s Server) {
	s.Router.HandleFunc("/tools", c.getAllTools).Methods("GET")
	s.Router.HandleFunc("/tools", c.createTool).Methods("POST")
}

func (c *ToolController) getAllTools(w http.ResponseWriter, r *http.Request) {
	response := c.repository.GetAllTools()
	json.NewEncoder(w).Encode(response)
}

func (c *ToolController) createTool(w http.ResponseWriter, r *http.Request) {
	var entity *entities.Tool
	json.NewDecoder(r.Body).Decode(&entity)
	entity, _ = c.repository.CreateTool(entity)
	json.NewEncoder(w).Encode(entity)
}
