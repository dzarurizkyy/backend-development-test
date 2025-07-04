package cmd

import (
	"api-test/config"
	"api-test/handler"
	"api-test/middleware"

	repository_implement "api-test/repository/implement"
	usecase_implement "api-test/usecase/implement"

	"github.com/labstack/echo/v4"
)

func Start() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	adminRepo := repository_implement.NewAdminRepository(db)
	authUsecase := usecase_implement.NewAuthUsecase(adminRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	terminalRepo := repository_implement.NewTerminalRepository(db)
	terminalUsecase := usecase_implement.NewTerminalUsecase(terminalRepo)
	terminalHandler := handler.NewTerminalHandler(terminalUsecase)

	e := echo.New()

	e.POST("/login", authHandler.Login)
	
	api := e.Group("/api")
	api.Use(middleware.JWTMiddleware())
	api.POST("/terminal", terminalHandler.CreateTerminal)

	e.Logger.Fatal(e.Start(":8080"))

}
