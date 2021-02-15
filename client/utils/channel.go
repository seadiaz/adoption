package utils

import "github.com/golang/glog"

// ReceiveResponses ...
func ReceiveResponses(channel chan string, quantity int) {
	counter := 0
	for value := range channel {
		glog.Infof(value)
		counter++
		if counter == quantity {
			close(channel)
		}
	}
}
