package teams

import (
	"fmt"

	"github.com/buger/goterm"
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/client/global"
	"github.com/seadiaz/adoption/client/utils"
)

func display(c *global.CommandHandler) {
	people := getTeams(c.BaseURL+path, c.APIKey)
	print(people)
}

func getTeams(url, apiKey string) []*Team {
	res, _ := utils.DoGetRequest(url, apiKey)
	output := make([]*Team, 0, 0)
	for _, item := range res {
		var person Team
		mapstructure.Decode(item, &person)
		output = append(output, &person)
	}
	return output
}

func print(people []*Team) {
	table := goterm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(table, "ID\tName\n")
	for _, p := range people {
		fmt.Fprintf(table, "%s\t%s\n", p.ID, p.Name)
	}
	goterm.Println(table)
	goterm.Flush()
}
