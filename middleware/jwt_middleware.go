package middleware

import (
	"api-test/config"
	"api-test/handler/dto"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			if auth == "" {
				return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
					Code:    http.StatusUnauthorized,
					Message: "missing authorization header",
				})
			}

			if !strings.HasPrefix(auth, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
					Code:    http.StatusUnauthorized,
					Message: "invalid authorization header format",
				})
			}

			tokenString := strings.TrimPrefix(auth, "Bearer ")

			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(config.AppConfig.JWT.Secret), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
					Code:    http.StatusUnauthorized,
					Message: "invalid token",
				})
			}

			c.Set("token", token)
			c.Set("claims", claims)

			c.Set("token", token)
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				c.Set("claims", claims)
			}

			return next(c)
		}
	}
}
