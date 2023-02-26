package route

import (
	handler "authService/internal/delivery/http"
	"authService/internal/usecase"

	"github.com/labstack/echo/v4"
)

func CreateRoutes(userUC usecase.User) *echo.Echo {
	e := echo.New()

	signUpHandler := handler.New(userUC)

	e.POST("/sign-up", signUpHandler.CreateUser)

	return e
}
