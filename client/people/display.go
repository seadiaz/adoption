package people

import (
	"fmt"

	tm "github.com/buger/goterm"
)

func displayV2(r *Repository) {
	people := r.GetPeople()
	print(people)
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
