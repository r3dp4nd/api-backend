package domain

import "context"

type CustomerRepository interface {
	Save(ctx context.Context, customer *Customer) error
	FindAll(ctx context.Context) ([]*Customer, error)
	Find(ctx context.Context, dni string) (*Customer, error)
	Update(ctx context.Context, customer *Customer) (*Customer, error)
}
