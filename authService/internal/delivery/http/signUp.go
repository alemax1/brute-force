package http

import (
	"authService/internal/delivery"
	"authService/internal/delivery/model"
	"authService/internal/domain"
	"authService/internal/usecase"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type signUp struct {
	userUC usecase.User
}

func New(userUC usecase.User) *signUp {
	return &signUp{
		userUC: userUC,
	}
}

func (s signUp) CreateUser(c echo.Context) error {
	var user domain.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: delivery.InvalidDataError.Error()})
	}

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: delivery.InvalidDataError.Error()})
	}

	if err := s.userUC.CreateUser(c.Request().Context(), user); err != nil {
		log.Err(err).Msg("error trying create user")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: delivery.ServerError.Error()})
	}

	return c.JSON(http.StatusCreated, model.SuccessfullResponse{Message: "user successfully created"})
}
