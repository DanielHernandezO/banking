package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DanielHernandezO/banking/internal/business/usecase"
	"github.com/gorilla/mux"
)

type CustomerHandler interface {
	GetAllCutomers(w http.ResponseWriter, r *http.Request)
	GetCustomer(w http.ResponseWriter, r *http.Request)
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
	Customer, usecaseErr := c.customerUsecase.FindAll()

	if usecaseErr != nil {
		w.WriteHeader(usecaseErr.Code)
		fmt.Fprintf(w, "Error: %v", usecaseErr.Message)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Customer)
	}
}

func (c *customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerIDStr := vars["customer_id"]

	customerID, err := strconv.Atoi(customerIDStr)
	if err != nil {
		http.Error(w, "Invalid customer_id", http.StatusBadRequest)
		return
	}

	customer, usecaseErr := c.customerUsecase.GetCustomer(customerID)
	if usecaseErr != nil {
		w.WriteHeader(usecaseErr.Code)
		fmt.Fprintf(w, "Error: %v", usecaseErr.Message)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
