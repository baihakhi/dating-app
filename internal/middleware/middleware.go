package middleware

import (
	"context"
	"net/http"

	"github.com/baihakhi/dating-app/internal/models"
	response "github.com/baihakhi/dating-app/internal/models/payload/responses"
	"github.com/labstack/echo/v4"
)

func SetMiddlewareAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := TokenValid(c.Request())
			if err != nil {
				c.JSON(http.StatusUnauthorized, response.MapResponse{
					Message: response.InvalidToken,
				})
				return err
			}

			if user.LastLogin == nil {
				err := response.ErrorBuilder("", response.AccessDenied)
				c.JSON(http.StatusForbidden, response.MapResponse{
					Message: err.Error(),
				})
				return err
			}
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), models.ACC, user)))
			return next(c)
		}
	}
}
