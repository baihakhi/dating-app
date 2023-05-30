package router

import (
	"net/http"

	"github.com/baihakhi/dating-app/internal/handler"
	"github.com/labstack/echo/v4"
)

func InitRouter(server *echo.Echo, handler *handler.Handler) {
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Internal API  App!")
	})

	v1 := server.Group("api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/register", handler.Register)
			user.POST("/login", handler.Login)
		}
	}
}
