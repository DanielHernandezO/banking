package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func mapRoutes(router *mux.Router, dependencies *Dependencies) {
	router.HandleFunc("/customers", dependencies.customerHandler.GetAllCutomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", dependencies.customerHandler.GetCustomer).Methods(http.MethodGet)
}
