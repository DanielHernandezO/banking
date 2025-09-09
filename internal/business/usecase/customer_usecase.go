package usecase

import (
	"github.com/DanielHernandezO/banking/internal/business/domain"
	"github.com/DanielHernandezO/banking/internal/business/gateway"
)

type CustomerUsecase interface {
	FindAll() ([]domain.Customer, *domain.AppError)
	GetCustomer(id int) (*domain.Customer, *domain.AppError)
}

type customerUsecase struct {
	customerRepository gateway.CustomerGateway
}

func NewCustomerUsecase(customerRepository gateway.CustomerGateway) *customerUsecase {
	return &customerUsecase{
		customerRepository: customerRepository,
	}
}

func (s *customerUsecase) FindAll() ([]domain.Customer, *domain.AppError) {
	return s.customerRepository.FindAll()
}

func (s *customerUsecase) GetCustomer(id int) (*domain.Customer, *domain.AppError) {
	return s.customerRepository.ById(id)
}
