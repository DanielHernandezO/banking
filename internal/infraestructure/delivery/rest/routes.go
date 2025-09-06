package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func mapRoutes(router *mux.Router, dependencies *Dependencies) {
	router.HandleFunc("/customers", dependencies.customerHandler.GetAllCutomers).Methods(http.MethodGet)
}
