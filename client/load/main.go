package load

import "github.com/golang/glog"

// CommandHandler ...
type CommandHandler struct {
}

// CommandHandlerParams ...
type CommandHandlerParams struct {
	Kind     string
	Filename string
}

// CreateCommandHandler ...
func CreateCommandHandler(baseUrl, apiKey string) *CommandHandler {
	return &CommandHandler{}
}

// Execute ...
func (*CommandHandler) Execute(params *CommandHandlerParams) error {
	glog.Infof("display command handler dispatched")
	// 	client := createClient(params.URL, params.Filename, params.APIKey)
	switch params.Kind {
	// 	case "adoptables":
	// 		client.LoadAdoptables()
	case "people":
		client.LoadPeople()
	// 	case "adoptions":
	// 		client.LoadAdoptions()
	// 	case "teams":
	// 		client.LoadTeams()
	// 	case "memberships":
	// 		client.LoadMemberships()
	// 	case "labels":
	// 		client.LoadLabels()
	default:
		glog.Fatalf("kind %s not supported", params.Kind)
	}
	return nil
}
