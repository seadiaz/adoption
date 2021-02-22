package teams

import (
	"fmt"

	"github.com/buger/goterm"
	"github.com/seadiaz/adoption/client/global"
)

func display(c *global.CommandHandler) {
	people := GetTeams(c.BaseURL+Path, c.APIKey)
	print(people)
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
