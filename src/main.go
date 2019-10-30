package main

import (
	"flag"
)

var (
	port   = flag.String("port", "3000", "listen port")
	server = flag.Bool("server", false, "run as a server")
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()
}

func main() {
	if *server {
		mainServer()
	}

	mainCLI()
}
