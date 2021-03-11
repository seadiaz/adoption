package people

import (
	"github.com/seadiaz/adoption/client/utils"
)

func load(r *Repository, filename string) {
	var people []Person
	utils.ReadCsvFile(filename, &people)
	for _, v := range people {
		r.SavePerson(&v)
	}
}
