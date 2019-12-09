package stepdefinitions

// World ...
type World struct {
	Adoptions     map[string]interface{}
	People        map[string]interface{}
	Adoptables    map[string]interface{}
	Teams         map[string]interface{}
	AdoptableList []interface{}
	PeopleList    []interface{}
	LabelList     []interface{}
}

// CreateWorld ...
func CreateWorld() *World {
	return &World{
		Adoptions:     make(map[string]interface{}),
		People:        make(map[string]interface{}),
		Adoptables:    make(map[string]interface{}),
		Teams:         make(map[string]interface{}),
		AdoptableList: make([]interface{}, 0, 0),
		PeopleList:    make([]interface{}, 0, 0),
		LabelList:     make([]interface{}, 0, 0),
	}
}

// Clear ...
func (w *World) Clear() {
	w.Adoptions = make(map[string]interface{})
	w.People = make(map[string]interface{})
	w.Adoptables = make(map[string]interface{})
	w.Teams = make(map[string]interface{})
	w.AdoptableList = make([]interface{}, 0, 0)
	w.PeopleList = make([]interface{}, 0, 0)
	w.LabelList = make([]interface{}, 0, 0)
}
