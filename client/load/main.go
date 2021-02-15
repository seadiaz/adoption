package load

import "github.com/golang/glog"

// CommandHandler ...
type CommandHandler struct {
	baseURL string
	apiKey  string
}

// CommandHandlerParams ...
type CommandHandlerParams struct {
	Kind     string
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
	// 	client := createClient(params.URL, params.Filename, params.APIKey)
	switch params.Kind {
	// 	case "adoptables":
	// 		client.LoadAdoptables()
	case "people":
		c.loadPeople(params.Filename)
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
