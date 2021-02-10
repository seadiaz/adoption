package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/golang/glog"
)

type apiClient struct {
	baseURL string
	apiKey  string
}

func createAPIClient(baseURL, apiKey string) *apiClient {
	return &apiClient{
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

func printBodyMessage(body io.Reader) {
	var output map[string]interface{}
	json.NewDecoder(body).Decode(&output)
	glog.Warning(output)
}

func (c *apiClient) doPostRequest(path string, body interface{}) error {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(body)
	url := c.baseURL + path
	req, _ := http.NewRequest(http.MethodPost, url, reqBodyBytes)
	req.Header.Add("Content-Type", "application/json")
	if c.apiKey != "" {
		req.Header.Add("Authorization", c.apiKey)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Error(err)
		return err
	}
	if res.StatusCode != 200 {
		glog.Errorf("request not succeed: %d", res.StatusCode)
		printBodyMessage(res.Body)
		return fmt.Errorf("do post request fail with status: %s", res.Status)
	}
	defer res.Body.Close()
	var output []map[string]interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return nil
}

func (c *apiClient) doGetRequest(path string) ([]map[string]interface{}, error) {
	url := c.baseURL + path
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Content-Type", "application/json")
	if c.apiKey != "" {
		req.Header.Add("Authorization", c.apiKey)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	if res.StatusCode != 200 {
		glog.Errorf("request not succeed: %s", res.Status)
		return nil, fmt.Errorf("do get request fail with status: %s", res.Status)
	}
	defer res.Body.Close()
	var output []map[string]interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output, nil
}
