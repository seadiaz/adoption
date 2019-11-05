package drivers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seadiaz/adoption/src/details"
	"github.com/seadiaz/adoption/src/details/adapters"
	usecases "github.com/seadiaz/adoption/src/details/adapters/usecases"
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

	toolRepository := adapters.CreateToolRepository(details.BuildMemoryPersistence())
	personRepository := adapters.CreatePersonRepository(details.BuildMemoryPersistence())
	teamRepository := adapters.CreateTeamRepository(details.BuildMemoryPersistence())

	toolService := usecases.CreateToolService(toolRepository)
	personService := usecases.CreatePersonService(personRepository)
	teamService := usecases.CreateTeamService(teamRepository)
	adoptionService := usecases.CreateAdoptionService(toolRepository, personRepository)

	toolController := adapters.CreateToolController(toolService, adoptionService)
	personController := adapters.CreatePersonController(personService)
	teamController := adapters.CreateTeamController(teamService)

	toolController.AddRoutes(server)
	personController.AddRoutes(server)
	teamController.AddRoutes(server)

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
