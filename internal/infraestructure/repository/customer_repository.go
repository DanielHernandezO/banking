package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/DanielHernandezO/banking/internal/business/domain"
	"github.com/DanielHernandezO/banking/internal/infraestructure/config"
)

type customerRepository struct {
	mysqlClient *sql.DB
}

func NewCustomerRepository() *customerRepository {
	return &customerRepository{
		mysqlClient: setUpMysqlDriver(),
	}
}

func setUpMysqlDriver() *sql.DB {
	databaseUser := config.GetStringPropetyBykey(config.DatabaseUser)
	databasePassword := config.GetStringPropetyBykey(config.DatabasePass)
	databaseHost := config.GetStringPropetyBykey(config.DatabaseHost)
	databasePort := config.GetStringPropetyBykey(config.DatabasePort)
	databaseSquema := config.GetStringPropetyBykey(config.DatabaseSquema)

	mysqlClient, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", databaseUser, databasePassword, databaseHost, databasePort, databaseSquema))
	if err != nil {
		panic(err)
	}

	mysqlClient.SetConnMaxLifetime(time.Minute * 3)
	mysqlClient.SetMaxOpenConns(10)
	mysqlClient.SetMaxIdleConns(10)
	return mysqlClient
}

func (c customerRepository) FindAll() ([]domain.Customer, *domain.AppError) {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
	rows, err := c.mysqlClient.Query(findAllSql)
	if err != nil {
		log.Println("Error querying customers:", err.Error())
		return nil, domain.NewUnexpectedError("unexpected database error")
	}

	var customers []domain.Customer
	for rows.Next() {
		var customer domain.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.ZipCode, &customer.DateOfBirth, &customer.Status)
		if err != nil {
			log.Println("Error scanning customer:", err.Error())
			return nil, domain.NewUnexpectedError("unexpected database error")
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c customerRepository) ById(id int) (*domain.Customer, *domain.AppError) {
	findByIdSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"
	row := c.mysqlClient.QueryRow(findByIdSql, id)

	var customer domain.Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.ZipCode, &customer.DateOfBirth, &customer.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.NewNotFoundError("customer not found")
		}
		log.Println("Error scanning customer by ID:", err.Error())
		return nil, domain.NewUnexpectedError("unexpected database error")
	}
	return &customer, nil
}
