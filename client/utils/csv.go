package utils

import (
	"fmt"
	"io/ioutil"

	"github.com/jszwec/csvutil"
)

// ReadCsvFile ...
func ReadCsvFile(filename string, output interface{}) {
	content, _ := ioutil.ReadFile(filename)
	if err := csvutil.Unmarshal(content, output); err != nil {
		fmt.Println("error:", err)
	}
}
