package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"../handlers"
)

// Route defines the Mux 
// router individual route 
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Routes is a collection 
// of individual Routes
var Routes = []Route{
	Route{"/", "GET", Index},
	Route{"/health", "GET", handlers.Health},

	Route{"/store/set", "POST", handlers.SetValue},
	Route{"/store/get/{key}", "GET", handlers.GetValue},

	Route{"/neighbor", "GET", handlers.GetNeighbor},
	Route{"/neighbor/{neighbor}", "GET", handlers.UpdateNeighbor},
}

// Index is the handler for the path "/"
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World GoServer\n")
}

// Logger is the middleware to
// log the incoming request
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"path":   r.URL,
			"method": r.Method,
		}).Info("incoming request")

		next.ServeHTTP(w, r)
	})
}

// Router returns a mux router
func Router() *mux.Router {
	router := mux.NewRouter()

	for _, route := range Routes {
		router.HandleFunc(
			route.Path,
			route.Handler,
		).Methods(route.Method)
	}

	router.Use(Logger)

	return router
}
