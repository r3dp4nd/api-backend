package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/r3dp4nd/api-backend/customer/internal/domain"
	"github.com/r3dp4nd/api-backend/customer/internal/dtos"
	"time"
)

type Application struct {
	customerRepository domain.CustomerRepository
	cityRepository     domain.CityRepository
}

func NewApplication(customerRepository domain.CustomerRepository, cityRepository domain.CityRepository) Application {
	return Application{customerRepository: customerRepository, cityRepository: cityRepository}
}

func (a Application) RegisterCustomer(ctx context.Context, customer dtos.RegisterCustomer) error {
	registerCustomer, err := domain.RegisterCustomer(customer.DNI, customer.Name, customer.LastName, customer.Telephone,
		customer.Email, customer.BirthDate, customer.City)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = a.customerRepository.Save(ctx, registerCustomer)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (a Application) GetAllCustomers(ctx context.Context) ([]*domain.Customer, error) {
	customers, err := a.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return customers, err
}

func (a Application) GetCustomer(ctx context.Context, getCustomer dtos.GetCustomer) (*domain.Customer, error) {
	customer, err := a.customerRepository.Find(ctx, getCustomer.DNI)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (a Application) UpdateCustomer(ctx context.Context, customerDNI string, updateCustomer dtos.UpdateCustomer) error {
	customer, err := a.customerRepository.Find(ctx, customerDNI)
	if err != nil {
		return err
	}

	propsMap, err := customerPropsMap(customer, updateCustomer)
	if err != nil {
		return err
	}
	a.customerRepository.Update(ctx, propsMap)

	return nil
}

func (a Application) DeleteCustomer(ctx context.Context, deleteCustomer dtos.DeleteCustomer) error {
	customer, err := a.customerRepository.Find(ctx, deleteCustomer.DNI)
	if err != nil {
		return err
	}

	if err = customer.Disable(); err != nil {
		return err
	}

	if _, err = a.customerRepository.Update(ctx, customer); err != nil {
		return err
	}

	return nil
}

func (a Application) GetCities(ctx context.Context) ([]*domain.City, error) {
	cities, err := a.cityRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return cities, err
}
func customerPropsMap(customer *domain.Customer, updateCustomer dtos.UpdateCustomer) (*domain.Customer, error) {
	if updateCustomer.DNI != "" {
		customer.DNI = updateCustomer.DNI
	}

	if updateCustomer.Name != "" {
		customer.Name = updateCustomer.Name
	}

	if updateCustomer.LastName != "" {
		customer.LastName = updateCustomer.LastName
	}

	if updateCustomer.Telephone != "" {
		customer.Telephone = updateCustomer.Telephone
	}

	if updateCustomer.Email != "" {
		customer.Email = updateCustomer.Email
	}

	if updateCustomer.BirthDate != "" {
		parse, err := time.Parse(time.DateOnly, updateCustomer.BirthDate)
		if err != nil {
			return nil, errors.New("format incorrect")
		}
		customer.BirthDate = &parse

	}

	if updateCustomer.City != "" {
		customer.City = updateCustomer.City
	}

	return customer, nil
}
