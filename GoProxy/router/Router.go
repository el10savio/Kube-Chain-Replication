package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"../handlers"
)

// Route ...
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Routes ...
var Routes = []Route{
	Route{"/", "GET", Index},

	Route{"/store/set", "POST", handlers.SetValue},
	Route{"/store/get/{key}", "GET", handlers.GetValue},

	Route{"/chain", "GET", handlers.GetPods},
	Route{"/chain/index", "GET", handlers.GetHeadTail},
	Route{"/chain/health", "GET", handlers.ChainHealth},
	Route{"/chain/neighbors", "GET", handlers.UpdateChainNeighbors},
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

// Logger ...
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"path":   r.URL,
			"method": r.Method,
		}).Info("incoming request")

		next.ServeHTTP(w, r)
	})
}

// Router ...
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
