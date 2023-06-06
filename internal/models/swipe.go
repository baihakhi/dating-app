package models

import (
	"time"

	"github.com/labstack/echo/v4"
)

type (
	Swipe struct {
		ID      uint64 `json:"swipe_id"`
		Swiper  uint64 `json:"swiper_id"`
		Swiped  uint64 `json:"swiped_id"`
		IsLiked bool   `json:"is_liked"`

		CreatedAt *time.Time `json:"created_at,omitempty"`
	}
)

func (s *Swipe) GetDataFromHTTPRequest(c echo.Context) error {
	if err := c.Bind(s); err != nil {
		return err
	}

	return nil
}
