package handler

import (
	"api-test/handler/dto"
	usecase_interface "api-test/usecase/interface"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TerminalHandler struct {
	Usecase usecase_interface.TerminalUsecase
}

func NewTerminalHandler(u usecase_interface.TerminalUsecase) *TerminalHandler {
	return &TerminalHandler{u}
}

func (t *TerminalHandler) CreateTerminal(c echo.Context) error {
	var req dto.CreateTerminalRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.CreateTerminalResponse{Code: http.StatusBadRequest, Message: "invalid request"})
	}

	if err := t.Usecase.CreateTerminal(req.Name); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CreateTerminalResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.CreateTerminalResponse{Code: http.StatusOK, Message: "success"})
}
