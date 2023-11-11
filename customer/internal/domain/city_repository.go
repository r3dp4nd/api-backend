package domain

import "context"

type CityRepository interface {
	FindAll(ctx context.Context) ([]*City, error)
}
