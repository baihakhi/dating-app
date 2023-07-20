package router

import (
	"net/http"

	"github.com/baihakhi/dating-app/internal/handler"
	"github.com/baihakhi/dating-app/internal/middleware"
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
			user.POST("/logout", handler.Logout, middleware.SetMiddlewareAuthentication())
			user.PUT("/verify/:username", handler.VerifyUser, middleware.SetMiddlewareAuthentication())
			user.GET("/:username", handler.GetUserDetail, middleware.SetMiddlewareAuthentication())
			user.GET("/:username", handler.GetUserDetail, middleware.SetMiddlewareAuthentication())
		}

		swipe := v1.Group("/swipe")
		{
			swipe.POST("/next", handler.SwipeAct, middleware.SetMiddlewareAuthentication())
		}
	}
}
