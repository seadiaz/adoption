package adoptables

import (
	"github.com/golang/glog"
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/client/utils"
)

const path = "/adoptables"

// Repository ...
type Repository struct {
	Client *utils.APIClient
}

// CreateRepository ...
func CreateRepository(client *utils.APIClient) *Repository {
	return &Repository{
		Client: client,
	}
}

// GetAdoptables ...
func (r *Repository) GetAdoptables() []*Adoptable {
	res, _ := r.Client.DoGetRequest(path)
	output := make([]*Adoptable, 0, 0)
	for _, item := range res {
		var person Adoptable
		mapstructure.Decode(item, &person)
		output = append(output, &person)
	}
	return output
}

// SaveAdoptable ...
func (r *Repository) SaveAdoptable(a *Adoptable) *Adoptable {
	err := r.Client.DoPostRequest(path, a)
	if err != nil {
		glog.Errorf("failt to save person %s", a.Name)
	}
	return a
}
