package people

import (
	"fmt"

	"github.com/seadiaz/adoption/client/global"

	tm "github.com/buger/goterm"
)

func display(c *global.CommandHandler) {
	people := GetPeople(c.BaseURL+Path, c.APIKey)
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

// func findPersonByEmail(people []*Person, email string) *Person {
// 	for _, item := range people {
// 		if item.Email == email {
// 			return item
// 		}
// 	}

// 	return nil
// }
