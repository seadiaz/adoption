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

// CommandHandler ...
type CommandHandler struct {
	baseURL string
	apiKey  string
}

// CommandHandlerParams ...
type CommandHandlerParams struct {
	Action   global.ActionType
	Filename string
}

// CreateCommandHandler ...
func CreateCommandHandler(baseURL, apiKey string) *CommandHandler {
	return &CommandHandler{
		baseURL,
		apiKey,
	}
}

// Execute ...
func (c *CommandHandler) Execute(params *CommandHandlerParams) error {
	glog.Infof("display command handler dispatched")
	switch params.Action {
	case global.Display:
		c.display()
	case global.Load:
		c.load(params.Filename)
	default:
		glog.Fatalf("action %s not supported", params.Action)
	}
	return nil
}
