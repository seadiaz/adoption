package adoptables

import (
	"github.com/golang/glog"
	"github.com/seadiaz/adoption/client/global"
)

// Adoptable ...
type Adoptable struct {
	ID       string `json:"id" mapstructure:"id"`
	Name     string `json:"name" csv:"Name" mapstructure:"name"`
	Strategy string `json:"strategy" csv:"Strategy" mapstructure:"strategy"`
}

// Execute ...
func Execute(r *Repository, params *global.CommandHandlerParams) error {
	if params.Kind != global.Adoptables {
		return nil
	}
	switch params.Action {
	case global.Display:
		display(r)
	case global.Load:
		load(r, params.Filename)
	default:
		glog.Fatalf("action %s not supported", params.Action)
	}
	return nil
}
