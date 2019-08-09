package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/seadiaz/adoption/src/details"
	"github.com/seadiaz/adoption/src/details/adapters"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()
}

func main() {
	glog.Info("server starting...")
	router := mux.NewRouter().StrictSlash(true)
	routerWrapper := &routerWrapper{router: router}
	httpServer := &http.Server{Addr: ":10000", Handler: router}
	server := adapters.CreateServer(httpServer, routerWrapper)

	memory := make(map[string]interface{})
	persistence := details.BuildMemoryPersistence(memory)
	toolRepository := adapters.CreateToolRepository(persistence)
	toolController := adapters.CreateToolController(toolRepository)
	toolController.AddRoutes(server)
	server.Run()
}

type routerWrapper struct {
	router *mux.Router
}

func (r *routerWrapper) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) adapters.Route {
	return &routeWrapper{route: r.router.HandleFunc(path, f)}
}

type routeWrapper struct {
	route *mux.Route
}

func (r *routeWrapper) Methods(methods ...string) adapters.Route {
	r.route.Methods(methods[0])
	return r
}
