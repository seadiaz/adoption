package display

import "github.com/golang/glog"

// CommandHandler ...
type CommandHandler struct {
}

// CommandHandlerParams ...
type CommandHandlerParams struct {
	Kind string
}

// CreateCommandHandler ...
func CreateCommandHandler(baseUrl, apiKey string) *CommandHandler {
	return &CommandHandler{}
}

// Execute ...
func (*CommandHandler) Execute(params *CommandHandlerParams) error {
	glog.Infof("display command handler dispatched")
	return nil
}
