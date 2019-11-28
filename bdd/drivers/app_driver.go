package drivers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/golang/glog"
)

// CreateToolWithName ...
func CreateToolWithName(name string) (map[string]interface{}, error) {
	path := "/tools"
	body := make(map[string]interface{})
	body["name"] = name
	res, err := postMessage(body, path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("create tool with name failed")
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
		glog.Error(res)
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

// AdoptToolByPerson ...
func AdoptToolByPerson(toolID string, personID string) (map[string]interface{}, error) {
	path := "/people/" + personID + "/tools"
	body := make(map[string]interface{})
	body["ID"] = toolID
	res, err := postMessage(body, path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("adopt tool by person failed")
	}
	defer res.Body.Close()
	var output map[string]interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output, nil
}

// CalculateAdoptionForTool ...
func CalculateAdoptionForTool(toolID string) (map[string]interface{}, error) {
	path := "/tools/" + toolID + "/adoption"
	res, err := getMessage(path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("calculate adoption for tool failed")
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
		return nil, errors.New("get message status code unexpected")
	}

	return res, nil
}

// GetAllTools ...
func GetAllTools() ([]interface{}, error) {
	path := "/tools"
	res, err := getMessage(path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("get all tools failed")
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

// AddLabelToTool ...
func AddLabelToTool(labelKind string, labelValue string, toolID string) ([]interface{}, error) {
	path := "/tools/" + toolID + "/labels"
	body := make(map[string]interface{})
	body["kind"] = labelKind
	body["value"] = labelValue
	res, err := postMessage(body, path)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("add label to tool failed")
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
