package teams

import (
	"fmt"

	"github.com/buger/goterm"
)

func display(r *Repository) {
	people := r.GetTeams()
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
