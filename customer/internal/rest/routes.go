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
				"data": customers,
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
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
				}

				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
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
				"data": "",
			})
		})

		group.PUT("/customer/:dni", func(ctx *gin.Context) {
			customers, err := application.GetAllCustomers(ctx.Request.Context())
			if err != nil {
				return
			}

			ctx.JSON(http.StatusCreated, gin.H{
				"data": customers,
			})
		})

		group.DELETE("/customer/:dni", func(c *gin.Context) {
			dni := c.Param("dni")
			err := application.DeleteCustomer(ctx, struct{ DNI string }{DNI: dni})
			if err != nil {
				fmt.Println(err)
				return
			}

			c.JSON(http.StatusOK, struct {
			}{})
		})

		group.GET("/city", func(c *gin.Context) {
			customers, err := application.GetCities(ctx)
			if err != nil {
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"data": customers,
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
		return "This field is required"
	case "min":
		return "Should be less than " + fe.Param()
	case "max":
		return "Should be greater than " + fe.Param()
	case "email":
		return "no cumple con el formato " + fe.Param()
	}
	return "Unknown error"
}
