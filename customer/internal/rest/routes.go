package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/r3dp4nd/api-backend/customer/internal/application"
	"github.com/r3dp4nd/api-backend/customer/internal/dtos"
)

func RegisterRoutes(ctx context.Context, engine *gin.Engine, application application.Application) {
	rootPathV1 := "/api/v1"

	group := engine.Group(rootPathV1)
	{
		group.GET("/customer", func(c *gin.Context) {
			customers, err := application.GetAllCustomers(ctx)
			if err != nil {
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"errors": err,
				"data":   customers,
				"code":   http.StatusOK,
			})
		})

		group.POST("/customer", func(c *gin.Context) {
			registerCustomer := dtos.RegisterCustomer{}
			err := c.ShouldBindJSON(&registerCustomer)
			if err != nil {
				var ve validator.ValidationErrors
				if errors.As(err, &ve) {
					out := make([]ErrorMsg, len(ve))
					for i, fe := range ve {
						out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
					}
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"errors": out,
						"data":   struct{}{},
						"code":   http.StatusBadRequest,
					})
				}

				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"errors": err.Error(),
					"data":   struct{}{},
					"code":   http.StatusBadRequest},
				)
				return
			}

			err = application.RegisterCustomer(ctx, registerCustomer)
			if err != nil {
				fmt.Println(err)
				return
			}
			if err != nil {
				return
			}

			c.JSON(http.StatusCreated, gin.H{
				"errors": err,
				"data":   struct{}{},
				"code":   http.StatusCreated,
			})
		})

		group.PUT("/customer/:dni", func(c *gin.Context) {
			dni := c.Param("dni")
			updateCustomer := dtos.UpdateCustomer{}
			err := c.ShouldBindJSON(&updateCustomer)
			if err != nil {
				var ve validator.ValidationErrors
				if errors.As(err, &ve) {
					out := make([]ErrorMsg, len(ve))
					for i, fe := range ve {
						out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
					}
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
				}

				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
				return
			}
			application.UpdateCustomer(ctx, dni, updateCustomer)
			if err != nil {
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"errors": err,
				"data":   struct{}{},
				"code":   http.StatusOK,
			})
		})

		group.DELETE("/customer/:dni", func(c *gin.Context) {
			dni := c.Param("dni")
			err := application.DeleteCustomer(ctx, struct{ DNI string }{DNI: dni})
			if err != nil {
				fmt.Println(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"errors": err,
				"data":   struct{}{},
				"code":   http.StatusOK,
			})
		})

		group.GET("/city", func(c *gin.Context) {
			cities, err := application.GetCities(ctx)
			if err != nil {
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"errors": err,
				"data":   cities,
				"code":   http.StatusOK,
			})
		})
	}

}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "El campo es requerido"
	case "min":
		return "El campo no la longitud necesaria de " + fe.Param()
	case "max":
		return "El campo no la longitud necesaria de " + fe.Param()
	case "email":
		return "El email es invalido" + fe.Param()
	}
	return "Unknown error"
}
