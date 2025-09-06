package rest

import (
	"github.com/DanielHernandezO/banking/internal/business/usecase"
	"github.com/DanielHernandezO/banking/internal/infraestructure/delivery/rest/handler"
)

type Dependencies struct {
	customerHandler handler.CustomerHandler
}

func buildDependencies() *Dependencies {
	//usecases
	customerUsecase := usecase.NewCustomerUsecase()

	return &Dependencies{
		customerHandler: handler.NewCustomerHandler(customerUsecase),
	}
}
