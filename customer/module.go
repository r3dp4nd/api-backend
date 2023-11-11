package customer

import (
	"context"
	"fmt"
	"github.com/r3dp4nd/api-backend/customer/internal/rest"

	"github.com/r3dp4nd/api-backend/customer/internal/application"
	"github.com/r3dp4nd/api-backend/customer/internal/mysql"
	"github.com/r3dp4nd/api-backend/internal/config"
)

type Module struct {
}

func (m *Module) StartUp(ctx context.Context, config config.AppConfig) error {
	fmt.Println(config)
	customerRepository := mysql.NewCustomerRepository(config.DB())
	cityRepository := mysql.NewCityRepository(config.DB())

	app := application.NewApplication(customerRepository, cityRepository)

	rest.RegisterRoutes(config.Engine(), app)

	return nil
}
