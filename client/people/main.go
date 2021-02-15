package people

import (
	"github.com/golang/glog"
	"github.com/seadiaz/adoption/client/global"
)

const peoplePath = "/people"

// Person ...
type Person struct {
	ID    string `json:"id" mapstructure:"id"`
	Name  string `json:"name" csv:"Name" mapstructure:"name"`
	Email string `json:"email" csv:"Email" mapstructure:"email"`
}

// Execute ...
func Execute(c *global.CommandHandler, params *global.CommandHandlerParams) error {
	if params.Kind != global.People {
		return nil
	}
	switch params.Action {
	case global.Display:
		display(c)
	case global.Load:
		load(c, params.Filename)
	default:
		glog.Fatalf("action %s not supported", params.Action)
	}
	return nil
}
