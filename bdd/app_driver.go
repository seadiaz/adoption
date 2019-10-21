package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seadiaz/adoption/src/details"
	"github.com/seadiaz/adoption/src/details/adapters"
	usecases "github.com/seadiaz/adoption/src/details/adapters/use_cases"
)

func start() {
	router := mux.NewRouter().StrictSlash(true)
	routerWrapper := &routerWrapper{router: router}
	httpServer := &http.Server{Addr: ":10000", Handler: router}
	server := adapters.CreateServer(httpServer, routerWrapper)

	toolRepository := adapters.CreateToolRepository(details.BuildMemoryPersistence())
	toolService := usecases.CreateToolService(toolRepository)
	toolController := adapters.CreateToolController(toolService)
	toolController.AddRoutes(server)

	personRepository := adapters.CreatePersonRepository(details.BuildMemoryPersistence())
	personService := usecases.CreatePersonService(personRepository)
	personController := adapters.CreatePersonController(personService)
	personController.AddRoutes(server)

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
