package people

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/client/global"
	"github.com/seadiaz/adoption/client/utils"

	tm "github.com/buger/goterm"
)

func display(c *global.CommandHandler) {
	people := getPerson(c.BaseURL+peoplePath, c.APIKey)
	print(people)
}

func getPerson(url, apiKey string) []*Person {
	res, _ := utils.DoGetRequest(url, apiKey)
	output := make([]*Person, 0, 0)
	for _, item := range res {
		var person Person
		mapstructure.Decode(item, &person)
		output = append(output, &person)
	}
	return output
}

func print(people []*Person) {
	table := tm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(table, "ID\tName\tEmail\n")
	for _, p := range people {
		fmt.Fprintf(table, "%s\t%s\t%s\n", p.ID, p.Name, p.Email)
	}
	tm.Println(table)
	tm.Flush()
}

// func findPersonByEmail(people []*Person, email string) *Person {
// 	for _, item := range people {
// 		if item.Email == email {
// 			return item
// 		}
// 	}

// 	return nil
// }
