package client

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
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
		log.Fatalln("couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	output := make([][]string, 0, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, record)
	}

	return output
}

func doPostRequest(body interface{}, url string, apiKey string) {
	glog.Info(url)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(body)
	req, _ := http.NewRequest(http.MethodPost, url, reqBodyBytes)
	req.Header.Add("Content-Type", "application/json")
	if apiKey != "" {
		req.Header.Add("Authorization", apiKey)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		glog.Error(err)
	}
	defer res.Body.Close()
	var output []map[string]interface{}
	json.NewDecoder(res.Body).Decode(&output)
	glog.Info("response:", output)
}

func doGetRequest(url string, apiKey string) []map[string]interface{} {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		glog.Error(err)
	}
	defer res.Body.Close()
	var output []map[string]interface{}
	json.NewDecoder(res.Body).Decode(&output)
	return output
}
