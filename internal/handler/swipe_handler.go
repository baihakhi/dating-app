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
			Data: 0,
		})
	}

	user, err := h.service.GetOneUserByUsername(account.Username)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.MapResponse{
			Message: response.ErrorBuilder(models.UUname, response.UNKN).Error(),
			Data: 0,
		})
	}

	if user.LastLogin == nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.ErrorBuilder(string(models.ACC), response.ULGN).Error(),
			Data: 0,
		})
	}
	data.Swiper = user.ID

	if data.Swiped > 0 {
		swipeLog, _ := h.service.GetSwipe(data.Swiped, user.ID)
		if swipeLog != nil {
			if swipeLog.IsLiked {
				matchID, err = h.service.CreateMatch(data.Swiped, user.ID)
				if err != nil {
					return c.JSON(http.StatusBadRequest, response.MapResponse{
						Message: err.Error(),
						Data: 0,
					})
				}
				if err := h.service.DeleteSwipe(swipeLog.ID); err != nil {
					return c.JSON(http.StatusBadRequest, response.MapResponse{
						Message: err.Error(),
						Data: 0,
					})
				}
			}
		}

		swipeID, err = h.service.CreateSwipe(user.Username, data, user.LastLogin)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.MapResponse{
				Message: err.Error(),
				Data: swipeID,
			})
		}
	}

	nextUser, err := h.service.NextUser(user.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
			Data: nextUser,
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
