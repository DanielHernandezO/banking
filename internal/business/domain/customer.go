package domain

type Customer struct {
	Id          int    `json:"id" xml:"id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip_code" xml:"zip_code"`
	DateOfBirth string `json:"date_of_birth" xml:"date_of_birth"`
	Status      int    `json:"status" xml:"status"`
}
