package middleware

import (
	"fmt"
	"net/http"
	jwtToken "pt-xyz-multifinance/pkg/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type Result struct {
	Code    int         `json:"Code"`
	Data    interface{} `json:"Data"`
	Message string      `json:"Message"`
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("sfdfsdfds")
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "unauthorized",
			})
		}
		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "unauthorized",
			})
		}
		c.Set("userLogin", claims)
		return next(c)
	}
}