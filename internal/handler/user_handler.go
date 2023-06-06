package handler

import (
	"net/http"

	"github.com/baihakhi/dating-app/internal/models"
	response "github.com/baihakhi/dating-app/internal/models/payload/responses"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(c echo.Context) error {
	data := new(models.User)

	if err := data.GetDataFromHTTPRequest(c); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.BADREQUEST,
		})
	}
	if err := data.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}
	result, err := h.service.CreateUser(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.MapResponse{
		Message: response.CREATED,
		Data: map[string]string{
			"username": result,
		},
	})
}

func (h *Handler) Login(c echo.Context) error {
	data := new(models.User)

	if err := data.GetDataFromHTTPRequest(c); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.BADREQUEST,
		})
	}
	if err := data.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	result, err := h.service.Login(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.AccessDenied,
		})
	}

	return c.JSON(http.StatusOK, response.MapResponse{
		Message: response.SUCCESS,
		Data: map[string]string{
			"token": result,
		},
	})
}

func (h *Handler) VerifyUser(c echo.Context) error {
	account := c.Request().Context().Value(models.ACC).(*models.User)
	username := c.Param("username")

	if username != account.Username {
		return c.JSON(http.StatusForbidden, response.MapResponse{
			Message: response.AccessDenied,
		})
	}

	user, err := h.service.GetOneUserByUsername(account.Username)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.MapResponse{
			Message: response.ErrorBuilder(models.UUname, response.UNKN).Error(),
		})
	}

	if err := h.service.PatchUserVerified(user.ID); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	if err := h.service.RemoveSwipeLimit(user.Username); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.MapResponse{
		Message: response.SUCCESS,
		Data:    response.SUCCESS,
	})
}
