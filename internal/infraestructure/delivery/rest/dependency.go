package rest

import (
	"github.com/DanielHernandezO/banking/internal/business/usecase"
	"github.com/DanielHernandezO/banking/internal/infraestructure/delivery/rest/handler"
	"github.com/DanielHernandezO/banking/internal/infraestructure/repository"
	"github.com/DanielHernandezO/banking/internal/infraestructure/repository/log"

	_ "github.com/go-sql-driver/mysql"
)

type Dependencies struct {
	customerHandler handler.CustomerHandler
}

func buildDependencies() *Dependencies {
	//logger
	loggerCollector := log.NewLoggerCollector()

	//repositories
	customerRepository := repository.NewCustomerRepository(loggerCollector)

	//usecases
	customerUsecase := usecase.NewCustomerUsecase(customerRepository)

	return &Dependencies{
		customerHandler: handler.NewCustomerHandler(customerUsecase),
	}
}
