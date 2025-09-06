package handler

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/DanielHernandezO/banking/internal/business/usecase"
)

type CustomerHandler interface {
	GetAllCutomers(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
	customerUsecase usecase.CustomerUsecase
}

func NewCustomerHandler(customerUsecase usecase.CustomerUsecase) *customerHandler {
	return &customerHandler{
		customerUsecase: customerUsecase,
	}
}

func (c *customerHandler) GetAllCutomers(w http.ResponseWriter, r *http.Request) {
	Customer, _ := c.customerUsecase.FindAll()

	if r.Header.Get("content-type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(Customer)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Customer)
	}
}
