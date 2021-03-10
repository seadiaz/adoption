package people

import (
	"github.com/golang/glog"
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/client/utils"
)

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

// FindPersonByEmail ...
func (r *Repository) FindPersonByEmail(email string) *Person {
	items := r.GetPeople()
	item := filterPersonByEmail(items, email)
	return item
}

func filterPersonByEmail(items []*Person, email string) *Person {
	for _, v := range items {
		if v.Email == email {
			return v
		}
	}

	return nil
}

// GetPeople ...
func (r *Repository) GetPeople() []*Person {
	res, _ := r.Client.DoGetRequest(Path)
	output := make([]*Person, 0, 0)
	for _, item := range res {
		var person Person
		mapstructure.Decode(item, &person)
		output = append(output, &person)
	}
	return output
}

// SavePerson ...
func (r *Repository) SavePerson(p *Person) *Person {
	err := r.Client.DoPostRequest(Path, p)
	if err != nil {
		glog.Errorf("failt to save person %s", p.Name)
	}
	return p
}
