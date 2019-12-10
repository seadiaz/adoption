package drivers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seadiaz/adoption/src/server/details"
	"github.com/seadiaz/adoption/src/server/details/adapters"
	usecases "github.com/seadiaz/adoption/src/server/details/adapters/usecases"
)

var port string = "10000"
var baseURL string = "http://localhost:" + port

// App ...
type App struct {
	server adapters.Server
}

// CreateApp ...
func CreateApp() *App {
	router := mux.NewRouter().StrictSlash(true)
	routerWrapper := &routerWrapper{router: router}
	httpServer := &http.Server{Addr: ":" + port, Handler: router}
	server := adapters.CreateServer(httpServer, routerWrapper)

	adoptableRepository := adapters.CreateAdoptableRepository(details.BuildMemoryPersistence())
	personRepository := adapters.CreatePersonRepository(details.BuildMemoryPersistence())
	teamRepository := adapters.CreateTeamRepository(details.BuildMemoryPersistence())

	adoptableService := usecases.CreateAdoptableService(adoptableRepository)
	personService := usecases.CreatePersonService(personRepository, adoptableRepository)
	teamService := usecases.CreateTeamService(teamRepository, personRepository)
	adoptionService := usecases.CreateAdoptionService(adoptableRepository, personRepository, teamRepository)

	adoptableController := adapters.CreateAdoptableController(adoptableService, adoptionService)
	personController := adapters.CreatePersonController(personService)
	teamController := adapters.CreateTeamController(teamService)

	adoptableController.AddRoutes(server.Router)
	personController.AddRoutes(server.Router)
	teamController.AddRoutes(server.Router)

	return &App{
		server: server,
	}
}

// StartApp ...
func (a *App) StartApp() {
	go a.server.Run()
}

// StopApp ...
func (a *App) StopApp() {
	a.server.Close()
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
