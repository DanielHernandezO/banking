package usecase

import (
	"github.com/DanielHernandezO/banking/internal/business/domain"
	"github.com/DanielHernandezO/banking/internal/business/gateway"
	"github.com/DanielHernandezO/banking/internal/infraestructure/delivery/rest/dto"
)

type CustomerUsecase interface {
	FindAll() ([]dto.CustomerDTO, *domain.AppError)
	GetCustomer(id int) (*dto.CustomerDTO, *domain.AppError)
}

type customerUsecase struct {
	customerRepository gateway.CustomerGateway
}

func NewCustomerUsecase(customerRepository gateway.CustomerGateway) *customerUsecase {
	return &customerUsecase{
		customerRepository: customerRepository,
	}
}

func (s *customerUsecase) FindAll() ([]dto.CustomerDTO, *domain.AppError) {
	customers, err := s.customerRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var customerDTOs []dto.CustomerDTO

	for _, customer := range customers {
		customerDTOs = append(customerDTOs, *customer.ToDTO())
	}
	return customerDTOs, nil
}

func (s *customerUsecase) GetCustomer(id int) (*dto.CustomerDTO, *domain.AppError) {
	customer, err := s.customerRepository.ById(id)
	if err != nil {
		return nil, err
	}

	return customer.ToDTO(), nil
}
