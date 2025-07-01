package handler

import (
	"net/http"

	"api-test/handler/dto"

	usecase_interface "api-test/usecase/interface"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Usecase usecase_interface.AuthUsecase
}

func NewAuthHandler(u usecase_interface.AuthUsecase) *AuthHandler {
	return &AuthHandler{u}
}

func (a *AuthHandler) Login(c echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.LoginResponse{Code: http.StatusBadRequest, Message: "invalid request", Token: "null"})
	}

	token, err := a.Usecase.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.LoginResponse{Code: http.StatusUnauthorized, Message: err.Error(), Token: "null"})
	}

	return c.JSON(http.StatusOK, dto.LoginResponse{Code: http.StatusOK, Message: "success", Token: token})
}
