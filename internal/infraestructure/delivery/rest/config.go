package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DanielHernandezO/banking/internal/infraestructure/config"
	"github.com/DanielHernandezO/banking/internal/infraestructure/delivery"
	"github.com/gorilla/mux"
)

type rest struct{}

func NewRest() delivery.Strategy {
	return &rest{}
}

func (r *rest) Start() {
	config.LoadConfigs()
	dependencies := buildDependencies()
	router := mux.NewRouter()
	mapRoutes(router, dependencies)
	log.Println("Starting server on port", config.GetStringPropetyBykey(config.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.GetStringPropetyBykey(config.Port)), router)
	if err != nil {
		log.Fatal(err)
	}
}
