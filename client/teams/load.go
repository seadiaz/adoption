package teams

import (
	"github.com/seadiaz/adoption/client/utils"
)

func load(r *Repository, filename string) {
	var teams []Team
	utils.ReadCsvFile(filename, &teams)
	for _, v := range teams {
		r.SaveTeam(&v)
	}
}
