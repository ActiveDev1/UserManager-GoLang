package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var token *jwt.Token

		claims := jwt.MapClaims{}
		authorization := c.Request().Header.Get("Authorization")
		if authorization == "" {
			return c.NoContent(http.StatusUnauthorized)
		}

		token, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected Signing Method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JwtSecret")), nil
		})

		if err != nil || !token.Valid {
			return c.NoContent(http.StatusUnauthorized)
		}

		c.Set("userId", claims["id"])

		return next(c)
	}

}
