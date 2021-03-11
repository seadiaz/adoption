package memberships

import (
	"github.com/golang/glog"
	"github.com/seadiaz/adoption/client/global"
)

// MembershipInput ...
type MembershipInput struct {
	PersonEmail string `csv:"Email"`
	TeamName    string `csv:"Team"`
}

// MembershipOutput ...
type MembershipOutput struct {
	TeamName   string
	PersonName string
}

// Execute ...
func Execute(r *Repository, params *global.CommandHandlerParams) error {
	if params.Kind != global.Memberships {
		return nil
	}
	switch params.Action {
	case global.Display:
		display(r, params.Parent)
	case global.Load:
		load(r, params.Parent, params.Filename)
	default:
		glog.Fatalf("action %s not supported", params.Action)
	}
	return nil
}
