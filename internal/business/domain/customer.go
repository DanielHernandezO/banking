package domain

import "github.com/DanielHernandezO/banking/internal/infraestructure/delivery/rest/dto"

const (
	STATUS_ACTIVE   = "active"
	STATUS_INACTIVE = "inactive"
)

type Customer struct {
	Id          int `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      int
}

func (c *Customer) ToDTO() *dto.CustomerDTO {
	return &dto.CustomerDTO{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

func (c *Customer) statusAsText() string {
	if c.Status == 1 {
		return STATUS_ACTIVE
	}
	return STATUS_INACTIVE
}
