package load

// import (
// 	"bytes"
// 	"encoding/csv"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"

// 	"github.com/golang/glog"
// )

// func readCsvFile(fileName string) [][]string {
// 	csvfile, err := os.Open(fileName)
// 	if err != nil {
// 		glog.Fatalln("couldn't open the csv file", err)
// 	}

// 	r := csv.NewReader(csvfile)

// 	output := make([][]string, 0, 0)
// 	for {
// 		record, err := r.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			glog.Fatal(err)
// 		}
// 		output = append(output, record)
// 	}

// 	return output
// }

// func receiveResponses(channel chan string, quantity int) {
// 	counter := 0
// 	for value := range channel {
// 		glog.Infof(value)
// 		counter++
// 		if counter == quantity {
// 			close(channel)
// 		}
// 	}
// }
