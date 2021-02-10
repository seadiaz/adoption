package client

import "github.com/golang/glog"

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
