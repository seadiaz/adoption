package adoptables

import (
	"github.com/seadiaz/adoption/client/utils"
)

func load(r *Repository, filename string) {
	var values []Adoptable
	utils.ReadCsvFile(filename, &values)
	for _, v := range values {
		r.SaveAdoptable(&v)
	}
}
