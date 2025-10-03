package domain

type Customer struct {
	Id          int    `json:"id" db:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
	Status      int    `json:"status"`
}
