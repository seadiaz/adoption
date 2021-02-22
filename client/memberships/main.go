package memberships

import (
	"github.com/golang/glog"
	"github.com/seadiaz/adoption/client/global"
	"github.com/seadiaz/adoption/client/people"
	"github.com/seadiaz/adoption/client/teams"
)

// Membership ...
type Membership struct {
	PersonEmail string
	TeamName    string
	Team        *teams.Team
	Person      *people.Person
}

// Execute ...
func Execute(c *global.CommandHandler, params *global.CommandHandlerParams) error {
	if params.Kind != global.Memberships {
		return nil
	}
	switch params.Action {
	case global.Display:
		display(c)
	// case global.Load:
	// 	load(c, params.Filename)
	default:
		glog.Fatalf("action %s not supported", params.Action)
	}
	return nil
}
