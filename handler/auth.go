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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	token, err := a.Usecase.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.LoginResponse{Message: "login success", Token: token})
}
