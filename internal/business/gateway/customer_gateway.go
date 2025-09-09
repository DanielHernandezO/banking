package gateway

import "github.com/DanielHernandezO/banking/internal/business/domain"

type CustomerGateway interface {
	FindAll() ([]domain.Customer, *domain.AppError)
	ById(id int) (*domain.Customer, *domain.AppError)
}
