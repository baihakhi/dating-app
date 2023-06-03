package models

import (
	"time"

	response "github.com/baihakhi/dating-app/internal/models/payload/responses"
	"github.com/labstack/echo/v4"
)

type (
	User struct {
		ID         uint64     `json:"user_id"`
		Username   string     `form:"username"`
		Email      string     `json:"email"`
		Password   string     `json:"password,omitempty"`
		Fullname   string     `json:"full_name"`
		Gender     string     `json:"gender"`
		Preference string     `json:"preference"`
		City       string     `json:"city"`
		Interest   string     `json:"interests"`
		IsVerified bool       `json:"is_verified"`
		LastLogin  *time.Time `json:"last_login"`

		CreatedAt *time.Time `json:"created_at,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}
	AccStr string
)

const (
	ACC     AccStr = "account"
	USwipes string = "swiped_today"
	UUname  string = "Username"
	UPass   string = "Password"
)

func (u *User) GetDataFromHTTPRequest(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		return err
	}

	return nil
}

func (u *User) Validate() error {
	if u.Username == "" {
		return response.ErrorBuilder(UUname, response.MDTR)
	}

	if u.Password == "" {
		return response.ErrorBuilder(UPass, response.MDTR)
	}

	return nil
}
