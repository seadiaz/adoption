package server

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/seadiaz/adoption/src/server/details"
	"github.com/seadiaz/adoption/src/server/details/adapters"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases"
)

// Params ...
type Params struct {
	Port      string
	Storage   string
	RedisPort int
	RedisHost string
}

// Boot ...
func Boot(params *Params) {
	glog.Info("server starting on port: ", params.Port)
	router := mux.NewRouter().StrictSlash(true)
	routerWrapper := &routerWrapper{router: router}
	httpServer := &http.Server{Addr: ":" + params.Port, Handler: cors.Default().Handler(router)}
	server := adapters.CreateServer(httpServer, routerWrapper)

	persistence := details.BuildMemoryPersistence()
	if params.Storage == "redis" {
		glog.Info("using redis for storage")
		persistence = details.BuildRedisPersistence(params.RedisHost, params.RedisPort)
	}

	adoptableRepository := adapters.CreateAdoptableRepository(persistence)
	personRepository := adapters.CreatePersonRepository(persistence)
	teamRepository := adapters.CreateTeamRepository(persistence)

	adoptableService := usecases.CreateAdoptableService(adoptableRepository)
	personService := usecases.CreatePersonService(personRepository, adoptableRepository)
	teamService := usecases.CreateTeamService(teamRepository, personRepository)
	adoptionService := usecases.CreateAdoptionService(adoptableRepository, personRepository, teamRepository)

	healthController := adapters.CreateHealthController()
	adoptableController := adapters.CreateAdoptableController(adoptableService, adoptionService)
	personController := adapters.CreatePersonController(personService)
	teamController := adapters.CreateTeamController(teamService)

	healthController.AddRoutes(server)
	adoptableController.AddRoutes(server.Router)
	personController.AddRoutes(server)
	teamController.AddRoutes(server)

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
