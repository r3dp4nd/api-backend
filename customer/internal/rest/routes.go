package rest

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/r3dp4nd/api-backend/customer/internal/application"
	"github.com/r3dp4nd/api-backend/customer/internal/dtos"
	"net/http"
)

func RegisterRoutes(engine *gin.Engine, application application.Application) {
	rootCustomerPath := "/api/customer"
	rootCityPath := "/api/city"

	engine.GET(rootCustomerPath, func(ctx *gin.Context) {
		customers, err := application.GetAllCustomers(ctx.Request.Context())
		if err != nil {
			return
		}

		ctx.JSON(200, gin.H{
			"data": customers,
		})
	})

	engine.POST(rootCustomerPath, func(ctx *gin.Context) {
		registerCustomer := dtos.RegisterCustomer{}
		err := ctx.ShouldBindJSON(&registerCustomer)
		if err != nil {
			fmt.Printf("%+v", err.Error())
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				out := make([]ErrorMsg, len(ve))
				for i, fe := range ve {
					out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
				}
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
			}

			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}

		err = application.RegisterCustomer(ctx.Request.Context(), registerCustomer)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err != nil {
			return
		}

		ctx.JSON(200, gin.H{
			"data": "",
		})
	})

	engine.PUT(rootCustomerPath, func(ctx *gin.Context) {
		customers, err := application.GetAllCustomers(ctx.Request.Context())
		if err != nil {
			return
		}

		ctx.JSON(200, gin.H{
			"data": customers,
		})
	})

	engine.DELETE(fmt.Sprintf("%s/:dni", rootCustomerPath), func(ctx *gin.Context) {
		dni := ctx.Param("dni")
		err := application.DeleteCustomer(ctx.Request.Context(), struct{ DNI string }{DNI: dni})
		if err != nil {
			fmt.Println(err)
			return
		}

		ctx.JSON(200, struct {
		}{})
	})

	engine.GET(rootCityPath, func(ctx *gin.Context) {
		customers, err := application.GetCities(ctx.Request.Context())
		if err != nil {
			return
		}

		ctx.JSON(200, gin.H{
			"data": customers,
		})
	})
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
