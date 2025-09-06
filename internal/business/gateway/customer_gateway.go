package gateway

import "github.com/DanielHernandezO/banking/internal/business/domain"

type CustomerGateway interface {
	FindAll() ([]domain.Customer, error)
}
