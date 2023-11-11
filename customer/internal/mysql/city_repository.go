package mysql

import (
	"context"
	"database/sql"
	"github.com/r3dp4nd/api-backend/customer/internal/domain"
)

type CityRepository struct {
	db *sql.DB
}

func NewCityRepository(db *sql.DB) CityRepository {
	return CityRepository{db: db}
}

func (c CityRepository) FindAll(ctx context.Context) ([]*domain.City, error) {
	const query = `SELECT * FROM cities`
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cities []*domain.City

	for rows.Next() {
		var city = &domain.City{}
		rows.Scan(&city.ID, &city.Name)
		cities = append(cities, city)

	}

	return cities, nil
}
