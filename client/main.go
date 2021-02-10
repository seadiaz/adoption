package client

import (
	"github.com/fatih/structs"
	"github.com/golang/glog"
)

// Params ...
type Params struct {
	Filename string
	URL      string
	APIKey   string
	Kind     string
}

type command interface {
	Notify(map[string]interface{}) error
}

// CommandDispatcher ...
type CommandDispatcher struct {
	commands  []command
	apiClient *apiClient
}

// BuildCommandDispatcher ...
func BuildCommandDispatcher(url, key string) *CommandDispatcher {
	c := &CommandDispatcher{
		commands:  make([]command, 0),
		apiClient: createAPIClient(url, key),
	}

	return c
}

// Execute ...
func (c *CommandDispatcher) Execute(cmdName string, params *Params) {
	glog.Info("execute called")
	paramsMap := structs.Map(params)
	paramsMap["Command"] = cmdName
	for _, item := range c.commands {
		item.Notify(paramsMap)
	}
}

// func createClient(url string, filename string, apiKey string) *client {
// 	return &client{
// 		url:      url,
// 		filename: filename,
// 		apiKey:   apiKey,
// 	}
// }
