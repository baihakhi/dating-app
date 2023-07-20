package middleware

import (
	"context"
	"net/http"

	"github.com/baihakhi/dating-app/internal/models"
	response "github.com/baihakhi/dating-app/internal/models/payload/responses"
	"github.com/labstack/echo/v4"
)

// SetMiddlewareAuthentication returns an echo.MiddlewareFunc that handles authentication for incoming requests.
func SetMiddlewareAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Validate the JWT token extracted from the incoming request.
			user, err := TokenValid(c.Request())
			if err != nil {
				c.JSON(http.StatusUnauthorized, response.MapResponse{
					Message: response.InvalidToken,
				})
				return err
			}

			// Check if the user has logged in before by inspecting the LastLogin field.
			if user.LastLogin == nil {
				err := response.ErrorBuilder("", response.AccessDenied)
				c.JSON(http.StatusForbidden, response.MapResponse{
					Message: err.Error(),
				})
				return err
			}

			// Set the user model in the request's context to make it accessible in subsequent handlers.
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), models.ACC, user)))
			return next(c)
		}
	}
}

