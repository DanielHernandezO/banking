package main

import (
	"github.com/DanielHernandezO/banking/internal/infraestructure/delivery/rest"
)

func main() {
	deliveryStrategy := rest.NewRest()
	deliveryStrategy.Start()
}
