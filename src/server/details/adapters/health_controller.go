package adapters

import (
	"encoding/json"
	"net/http"
)

// HealthController ...
type HealthController struct {
}

// CreateHealthController ...
func CreateHealthController() *HealthController {
	return &HealthController{}
}

// AddRoutes ...
func (c *HealthController) AddRoutes(s Server) {
	s.Router.HandleFunc("/health", c.getStatus).Methods("GET")
}

func (c *HealthController) getStatus(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]string)
	res["status"] = "ok"
	res["message"] = "all up and running"
	w.Header().Add("Content-Type", "application/json")
	output := CreateHealthResponseMap(res)
	json.NewEncoder(w).Encode(output)
}
