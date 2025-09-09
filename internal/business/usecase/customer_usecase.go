package usecase

import (
	"github.com/DanielHernandezO/banking/internal/business/domain"
	"github.com/DanielHernandezO/banking/internal/business/gateway"
)

type CustomerUsecase interface {
	FindAll() ([]domain.Customer, error)
}

type customerUsecase struct {
	customerRepository gateway.CustomerGateway
}

func NewCustomerUsecase(customerRepository gateway.CustomerGateway) *customerUsecase {
	return &customerUsecase{
		customerRepository: customerRepository,
	}
}

func (s *customerUsecase) FindAll() ([]domain.Customer, error) {
	users, err := s.customerRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
