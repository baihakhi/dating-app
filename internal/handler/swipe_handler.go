package handler

import (
	"net/http"

	"github.com/baihakhi/dating-app/internal/models"
	response "github.com/baihakhi/dating-app/internal/models/payload/responses"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SwipeAct(c echo.Context) error {
	var swipeID, matchID int64
	data := new(models.Swipe)
	account := c.Request().Context().Value(models.ACC).(*models.User)

	if err := data.GetDataFromHTTPRequest(c); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.BADREQUEST,
		})
	}

	user, err := h.service.GetOneUserByUsername(account.Username)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.MapResponse{
			Message: response.ErrorBuilder(models.UUname, response.UNKN).Error(),
		})
	}

	if data.Swiped > 0 {
		swipeLog, _ := h.service.GetSwipe(data.Swiped, user.ID)
		if swipeLog != nil {
			if swipeLog.IsLiked {
				matchID, err = h.service.CreateMatch(data.Swiped, user.ID)
				if err != nil {
					return c.JSON(http.StatusBadRequest, response.MapResponse{
						Message: err.Error(),
					})
				}
				if err := h.service.DeleteSwipe(swipeLog.ID); err != nil {
					return c.JSON(http.StatusBadRequest, response.MapResponse{
						Message: err.Error(),
					})
				}
			}
		}

		swipeID, err = h.service.CreateSwipe(user.Username, user.ID, data.Swiped, data.IsLiked)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.MapResponse{
				Message: err.Error(),
			})
		}
	}

	nextUser, err := h.service.NextUser(user.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.MapResponse{
		Message: response.CREATED,
		Data: map[string]interface{}{
			"swipe_id": swipeID,
			"match_id": matchID,
			"next":     nextUser,
		},
	})
}
