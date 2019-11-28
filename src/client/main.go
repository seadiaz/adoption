package client

import (
	"github.com/golang/glog"
)

type client struct {
	url      string
	filename string
	apiKey   string
}

// Params ...
type Params struct {
	Filename string
	URL      string
	APIKey   string
	Kind     string
}

func createClient(url string, filename string, apiKey string) *client {
	return &client{
		url:      url,
		filename: filename,
		apiKey:   apiKey,
	}
}

// LoadData ...
func LoadData(params *Params) {
	client := createClient(params.URL, params.Filename, params.APIKey)
	switch params.Kind {
	case "tools":
		client.LoadTools()
	case "people":
		client.LoadPeople()
	case "adoptions":
		client.LoadAdoptions()
	case "teams":
		client.LoadTeams()
	case "memberships":
		client.LoadMemberships()
	case "labels":
		client.LoadLabels()
	default:
		glog.Fatalf("kind %s not supported", params.Kind)
	}
}
