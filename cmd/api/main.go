package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/r3dp4nd/api-backend/customer"
	"github.com/r3dp4nd/api-backend/internal/server"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/r3dp4nd/api-backend/internal/config"
)

type app struct {
	db     *sql.DB
	engine *gin.Engine
}

func (a *app) DB() *sql.DB {
	return a.db
}

func (a *app) Engine() *gin.Engine {
	return a.engine
}

func main() {
	envConfig := config.NewConfig()
	fmt.Printf("%+v", envConfig)

	appConfig := &app{}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", envConfig.DBUsername, envConfig.DBPassword, envConfig.DBName))
	if err != nil {
		return
	}

	appConfig.db = db

	srv := server.NewServer(envConfig.Port)

	appConfig.engine = srv.Engine()

	srv.Engine().GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	customerModule := &customer.Module{}
	customerModule.StartUp(context.Background(), appConfig)

	srv.Run()

}
