package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/r3dp4nd/api-backend/customer/internal/domain"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return CustomerRepository{db: db}
}

func (c CustomerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	const query = `INSERT INTO customers (dni, name, last_name, telephone, email, birthdate, city_id, enabled) 
		Values(?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := c.db.ExecContext(ctx, query, customer.DNI, customer.Name, customer.LastName, customer.Telephone,
		customer.Email, customer.BirthDate, customer.City, customer.Enabled)

	return err
}

func (c CustomerRepository) FindAll(ctx context.Context) ([]*domain.Customer, error) {
	const query = `SELECT   dni, name, last_name, telephone, email, birthdate,
         city_id  FROM customers WHERE enabled != 0`
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []*domain.Customer

	for rows.Next() {
		fmt.Printf("%+v", rows)
		var customer = &domain.Customer{}
		//rows.Scan(&customer.DNI, &customer.Name, &customer.LastName, &customer.Telephone, &customer.Email,
		//	&customer.BirthDate, &customer.City)
		rows.Scan(&customer.DNI, &customer.Name, &customer.LastName, &customer.Telephone, &customer.Email,
			&customer.BirthDate, &customer.City)

		customers = append(customers, customer)

	}

	fmt.Printf("%+v\n", customers)

	return customers, nil

}

func (c CustomerRepository) Find(ctx context.Context, dni string) (*domain.Customer, error) {
	fmt.Println(dni)
	const query = `SELECT dni, name, last_name, telephone, email, birthdate, city_id FROM customers WHERE dni = ? and enabled != 0 LIMIT 1`

	row := c.db.QueryRowContext(ctx, query, dni)
	var customer = &domain.Customer{}
	err := row.Scan(&customer.DNI, &customer.Name, &customer.LastName, &customer.Telephone, &customer.Email,
		&customer.BirthDate, &customer.City)
	return customer, err
}

func (c CustomerRepository) Update(ctx context.Context, customer *domain.Customer) (*domain.Customer, error) {
	fmt.Printf("%+v", customer)
	const query = `UPDATE customers SET name = ?, last_name = ?, telephone = ?, email = ?, birthdate = ?,
        city_id = ?, enabled = ? WHERE dni = ?`
	_, err := c.db.ExecContext(ctx, query, customer.Name, customer.LastName, customer.Telephone,
		customer.Email, customer.BirthDate, customer.City, customer.Enabled, customer.DNI)
	return customer, err
}
