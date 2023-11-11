package config

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type AppConfig interface {
	DB() *sql.DB
	Engine() *gin.Engine
}
