package stepdefinitions

// World ...
type World struct {
	Adoptions map[string]interface{}
	People    map[string]interface{}
	Tools     map[string]interface{}
	ToolList  []interface{}
}

// CreateWorld ...
func CreateWorld() *World {
	return &World{
		Adoptions: make(map[string]interface{}),
		People:    make(map[string]interface{}),
		Tools:     make(map[string]interface{}),
		ToolList:  make([]interface{}, 0, 0),
	}
}

// Clear ...
func (w *World) Clear() {
	w.Adoptions = make(map[string]interface{})
	w.People = make(map[string]interface{})
	w.Tools = make(map[string]interface{})
	w.ToolList = make([]interface{}, 0, 0)
}
