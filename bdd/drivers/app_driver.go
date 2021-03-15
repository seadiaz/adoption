package drivers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/golang/glog"
)

// CreateAdoptableWithName ...
func CreateAdoptableWithName(name string) (map[string]interface{}, error) {
	path := "/adoptables"
	body := make(map[string]interface{})
	body["name"] = name
	body["strategy"] = "single-member"
	res, err := postMessage(body, path)
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("create adoptable with name failed")
	}
	defer res.Body.Close()
	var output map[string]interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output, nil
}

func postMessage(body map[string]interface{}, path string) (*http.Response, error) {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(body)
	req, _ := http.NewRequest(http.MethodPost, baseURL+path, reqBodyBytes)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("post message failed")
	}
	if res.StatusCode < 200 || 300 <= res.StatusCode {
		glog.Warning(res)
		printBodyMessage(res.Body)
		return nil, errors.New("post message status code unexpected failed")
	}

	return res, nil
}

// CreatePersonWithName ...
func CreatePersonWithName(name string) (map[string]interface{}, error) {
	path := "/people"
	body := make(map[string]interface{})
	body["name"] = name
	body["email"] = name + "@dummy.tld"
	res, err := postMessage(body, path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("create person with name failed")
	}
	defer res.Body.Close()
	var output map[string]interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output, nil
}

// CreateTeamWithName ...
func CreateTeamWithName(name string) (map[string]interface{}, error) {
	path := "/teams"
	body := make(map[string]interface{})
	body["name"] = name
	res, err := postMessage(body, path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("create team with name failed")
	}
	defer res.Body.Close()
	var output map[string]interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output, nil
}

// AdoptAdoptableByPerson ...
func AdoptAdoptableByPerson(adoptableID string, personID string) (map[string]interface{}, error) {
	path := "/people/" + personID + "/adoptables"
	body := make(map[string]interface{})
	body["id"] = adoptableID
	res, err := postMessage(body, path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("adopt adoptable by person failed")
	}
	defer res.Body.Close()
	var output map[string]interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output, nil
}

// CalculateAdoptionForAdoptable ...
func CalculateAdoptionForAdoptable(adoptableID string) (map[string]interface{}, error) {
	path := "/adoptables/" + adoptableID + "/adoption"
	res, err := getMessage(path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("calculate adoption for adoptable failed")
	}
	defer res.Body.Close()
	var output map[string]interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output, nil
}

func getMessage(path string) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, baseURL+path, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("get message failed")
	}
	if res.StatusCode < 200 || 300 <= res.StatusCode {
		glog.Error(res)
		printBodyMessage(res.Body)
		return nil, errors.New("get message status code unexpected")
	}

	return res, nil
}

// GetAllAdoptables ...
func GetAllAdoptables() ([]interface{}, error) {
	path := "/adoptables"
	res, err := getMessage(path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("get all adoptables failed")
	}
	defer res.Body.Close()
	var output []interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output, nil
}

// GetAllAdoptablesFilterByLabel ...
func GetAllAdoptablesFilterByLabel(kind string, value string) ([]interface{}, error) {
	path := "/adoptables?labels=" + kind + ":" + value
	res, err := getMessage(path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("get all adoptables failed")
	}
	defer res.Body.Close()
	var output []interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output, nil
}

// GetAllPeople ...
func GetAllPeople() ([]interface{}, error) {
	path := "/people"
	res, err := getMessage(path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("get all people failed")
	}
	defer res.Body.Close()
	var output []interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output, nil
}

// AddMemberToTeam ...
func AddMemberToTeam(personID string, teamID string) ([]interface{}, error) {
	path := "/teams/" + teamID + "/people"
	body := make(map[string]interface{})
	body["id"] = personID
	res, err := postMessage(body, path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("add member to team failed")
	}
	defer res.Body.Close()
	var output []interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output, nil
}

// AddLabelToAdoptable ...
func AddLabelToAdoptable(labelKind string, labelValue string, adoptableID string) ([]interface{}, error) {
	path := "/adoptables/" + adoptableID + "/labels"
	body := make(map[string]interface{})
	body["kind"] = labelKind
	body["value"] = labelValue
	res, err := postMessage(body, path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("add label to adoptable failed")
	}
	defer res.Body.Close()
	var output []interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output, nil
}

// GetMembersFromTeam ...
func GetMembersFromTeam(id string) ([]interface{}, error) {
	path := "/teams/" + id + "/people"
	res, err := getMessage(path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("get members from team failed")
	}
	defer res.Body.Close()
	var output []interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output, nil
}

func printBodyMessage(body io.Reader) {
	var output map[string]interface{}
	json.NewDecoder(body).Decode(&output)
	glog.Warning(output)
}
