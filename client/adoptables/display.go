package adoptables

import (
	"fmt"

	tm "github.com/buger/goterm"
)

func display(r *Repository) {
	values := r.GetAdoptables()
	print(values)
}

func print(values []*Adoptable) {
	table := tm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(table, "ID\tName\tStrategy\n")
	for _, p := range values {
		fmt.Fprintf(table, "%s\t%s\t%s\n", p.ID, p.Name, p.Strategy)
	}
	tm.Println(table)
	tm.Flush()
}
