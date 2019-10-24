package stepdefinitions

import (
	"github.com/seadiaz/adoption/bdd/drivers"
)

var tools map[string]interface{} = make(map[string]interface{})

// ThereIsAToolNamed ...
func ThereIsAToolNamed(name string) error {
	res, err := drivers.CreateToolWithName(name)
	tools[name] = res
	return err
}
