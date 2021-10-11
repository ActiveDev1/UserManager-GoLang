package middleware

import (
	"UserManager/container"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasttemplate"
)

func InitLoggerMiddleware(e *echo.Echo, container container.Container) {
	e.Use(RequestLoggerMiddleware(container))
}

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

func RequestLoggerMiddleware(container container.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			if err := next(c); err != nil {
				c.Error(err)
			}

			template := fasttemplate.New(container.GetConfig().Log.RequestLogFormat, "${", "}")
			logstr := template.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
				switch tag {
				case "remote_ip":
					return w.Write([]byte(c.RealIP()))
				case "uri":
					return w.Write([]byte(req.RequestURI))
				case "method":
					return w.Write([]byte(req.Method))
				case "status":
					return w.Write([]byte(strconv.Itoa(res.Status)))
				default:
					return w.Write([]byte(""))
				}
			})
			container.GetLogger().GetZapLogger().Infof(logstr)
			return nil
		}
	}
}
