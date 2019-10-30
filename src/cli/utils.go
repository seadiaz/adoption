package cli

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

func doPostRequest(body interface{}, url string) {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(body)
	req, _ := http.NewRequest(http.MethodPost, url, reqBodyBytes)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		glog.Error(err)
	}
	defer res.Body.Close()
	glog.Infof("payload %s created", body)
}
