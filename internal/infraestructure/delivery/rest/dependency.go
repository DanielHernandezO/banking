package rest

import (
	"github.com/DanielHernandezO/banking/internal/business/usecase"
	"github.com/DanielHernandezO/banking/internal/infraestructure/delivery/rest/handler"
	"github.com/DanielHernandezO/banking/internal/infraestructure/repository"

	_ "github.com/go-sql-driver/mysql"
)

type Dependencies struct {
	customerHandler handler.CustomerHandler
}

func buildDependencies() *Dependencies {
	//repositories
	customerRepository := repository.NewCustomerRepository()

	//usecases
	customerUsecase := usecase.NewCustomerUsecase(customerRepository)

	return &Dependencies{
		customerHandler: handler.NewCustomerHandler(customerUsecase),
	}
}
