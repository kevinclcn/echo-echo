package server

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func setupLogin()  {
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return c.String(http.StatusOK, "Welcome " + name+ "!")
	})

	e.POST("/login", func(c echo.Context) error {

		user := &User{}
		if err := c.Bind(user); err != nil {
			return err
		}

		if user.UserName == "Long" && user.Password == "Cheng" {
			token := jwt.New(jwt.SigningMethodHS256)

			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = "Long"
			claims["admin"] = true
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, map[string]string {
				"token": t,
			})
		}

		return echo.ErrUnauthorized
	})
}