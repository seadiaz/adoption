package teams

import (
	"github.com/golang/glog"
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/client/people"
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

// FindTeamByName ...
func (r *Repository) FindTeamByName(teamName string) *Team {
	teams := r.GetTeams()
	team := filterTeamByName(teams, teamName)
	return team
}

func filterTeamByName(teams []*Team, name string) *Team {
	for _, v := range teams {
		if v.Name == name {
			return v
		}
	}

	return nil
}

// GetMembersByTeam ...
func (r *Repository) GetMembersByTeam(team *Team) []*people.Person {
	res, _ := r.Client.DoGetRequest(Path + "/" + team.ID + people.Path)
	output := make([]*people.Person, 0, 0)
	for _, item := range res {
		var person people.Person
		mapstructure.Decode(item, &person)
		output = append(output, &person)
	}
	return output
}

// AddMemberToTeam ...
func (r *Repository) AddMemberToTeam(p *people.Person, t *Team) error {
	err := r.Client.DoPostRequest(Path+"/"+t.ID+people.Path, p)
	if err != nil {
		glog.Errorf("saving %s to %s failed", p.Name, t.Name)
	}

	return nil
}

// GetTeams ...
func (r *Repository) GetTeams() []*Team {
	res, _ := r.Client.DoGetRequest(Path)
	output := make([]*Team, 0, 0)
	for _, item := range res {
		var person Team
		mapstructure.Decode(item, &person)
		output = append(output, &person)
	}
	return output
}

// SaveTeam ...
func (r *Repository) SaveTeam(t *Team) *Team {
	err := r.Client.DoPostRequest(Path, t)
	if err != nil {
		glog.Errorf("fail to save person %s", t.Name)
	}
	return t
}
