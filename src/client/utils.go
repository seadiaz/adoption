package client

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/golang/glog"
)

type apiClient struct {
	key string
}

func createAPIClient(key string) *apiClient {
	return &apiClient{
		key: key,
	}
}

func readCsvFile(fileName string) [][]string {
	csvfile, err := os.Open(fileName)
	if err != nil {
		glog.Fatalln("couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	output := make([][]string, 0, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			glog.Fatal(err)
		}
		output = append(output, record)
	}

	return output
}

func doPostRequest(body interface{}, url string, apiKey string) error {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(body)
	req, _ := http.NewRequest(http.MethodPost, url, reqBodyBytes)
	req.Header.Add("Content-Type", "application/json")
	if apiKey != "" {
		req.Header.Add("Authorization", apiKey)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Error(err)
		return err
	}
	if res.StatusCode != 200 {
		glog.Errorf("request not succeed: %d", res.StatusCode)
		return fmt.Errorf("do post request fail with status: %s", res.Status)
	}
	defer res.Body.Close()
	var output []map[string]interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return nil
}

func doGetRequest(url string, apiKey string) ([]map[string]interface{}, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Content-Type", "application/json")
	if apiKey != "" {
		req.Header.Add("Authorization", apiKey)
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

func receiveResponses(channel chan string, quantity int) {
	counter := 0
	for value := range channel {
		glog.Infof(value)
		counter++
		if counter == quantity {
			close(channel)
		}
	}
}
