package usecase

import "github.com/DanielHernandezO/banking/internal/business/domain"

type CustomerUsecase interface {
	FindAll() ([]domain.Customer, error)
}

type customerUsecase struct{}

func NewCustomerUsecase() *customerUsecase {
	return &customerUsecase{}
}

func (s *customerUsecase) FindAll() ([]domain.Customer, error) {
	return []domain.Customer{
		{Id: "1", Name: "John Doe", City: "New York", ZipCode: "10001", DateOfBirth: "1990-01-01", Status: "active"},
		{Id: "2", Name: "Jane Smith", City: "Los Angeles", ZipCode: "90001", DateOfBirth: "1985-05-15", Status: "inactive"},
		{Id: "3", Name: "Alice Johnson", City: "Chicago", ZipCode: "60601", DateOfBirth: "1992-07-20", Status: "active"},
	}, nil
}
