package domain

import (
	"errors"
	"github.com/r3dp4nd/api-backend/customer/internal/helper"
	"time"
)

var (
	ErrDNICannotBlank         = errors.New(`the customer DNI cannot be blank`)
	ErrCustomerIsNotOldEnough = errors.New("the customer is not old enough")
	ErrEmailIsInvalid         = errors.New("the customer email is invalid")
)

const minAge = 18

type Customer struct {
	DNI       string     `json:"dni"`
	Name      string     `json:"name"`
	LastName  string     `json:"lastName"`
	Telephone string     `json:"telephone"`
	Email     string     `json:"email"`
	BirthDate *time.Time `json:"birthDate" `
	City      string     `json:"city"`
	Enabled   bool       `json:"enabled,omitempty"`
}

func RegisterCustomer(dni, name, lastName, telephone string, email string, birthDate string, city string) (*Customer, error) {
	if dni == "" {
		return nil, ErrDNICannotBlank
	}

	parseTime, err := time.Parse(time.DateOnly, birthDate)
	if err != nil {
		return nil, errors.New("format incorrect")
	}

	if helper.IsValidAge(parseTime, minAge) {
		return nil, ErrCustomerIsNotOldEnough
	}

	if !helper.IsValidEmail(email) {
		return nil, ErrEmailIsInvalid
	}

	return &Customer{
		DNI:       dni,
		Name:      name,
		LastName:  lastName,
		Telephone: telephone,
		Email:     email,
		BirthDate: &parseTime,
		City:      city,
		Enabled:   true,
	}, nil
}

func (c *Customer) Enable() error {
	if c.Enabled {
		return nil
	}

	c.Enabled = true

	return nil
}

func (c *Customer) Disable() error {
	if !c.Enabled {
		return nil
	}

	c.Enabled = false

	return nil
}
