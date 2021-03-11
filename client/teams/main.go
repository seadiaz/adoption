package teams

import (
	"github.com/golang/glog"
	"github.com/seadiaz/adoption/client/global"
)

// Path ..
const Path = "/teams"

// Team ...
type Team struct {
	ID   string `json:"id,omitempty" mapstructure:"id"`
	Name string `json:"name" csv:"Name" mapstructure:"name"`
}

// Execute ...
func Execute(r *Repository, params *global.CommandHandlerParams) error {
	if params.Kind != global.Teams {
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
