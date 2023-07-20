// Code generated by mockery v2.14.0. DO NOT EDIT.

package repositoriesMock

import (
	models "github.com/baihakhi/dating-app/internal/models"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Repositories is an autogenerated mock type for the Repositories type
type Repositories struct {
	mock.Mock
}

// CreateMatch provides a mock function with given fields: user1, user2
func (_m *Repositories) CreateMatch(user1 uint64, user2 uint64) (int64, error) {
	ret := _m.Called(user1, user2)

	var r0 int64
	if rf, ok := ret.Get(0).(func(uint64, uint64) int64); ok {
		r0 = rf(user1, user2)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(user1, user2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSwipe provides a mock function with given fields: swiper, swiped, is_liked
func (_m *Repositories) CreateSwipe(swiper uint64, swiped uint64, is_liked bool) (int64, error) {
	ret := _m.Called(swiper, swiped, is_liked)

	var r0 int64
	if rf, ok := ret.Get(0).(func(uint64, uint64, bool) int64); ok {
		r0 = rf(swiper, swiped, is_liked)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64, bool) error); ok {
		r1 = rf(swiper, swiped, is_liked)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: data
func (_m *Repositories) CreateUser(data *models.User) (string, error) {
	ret := _m.Called(data)

	var r0 string
	if rf, ok := ret.Get(0).(func(*models.User) string); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.User) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSwipe provides a mock function with given fields: userID
func (_m *Repositories) DeleteSwipe(userID uint64) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetOneUsersByUsername provides a mock function with given fields: username
func (_m *Repositories) GetOneUsersByUsername(username string) (*models.User, error) {
	ret := _m.Called(username)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPasswordByUsername provides a mock function with given fields: username
func (_m *Repositories) GetPasswordByUsername(username string) (string, error) {
	ret := _m.Called(username)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSwipe provides a mock function with given fields: swiperID, userID
func (_m *Repositories) GetSwipe(swiperID uint64, userID uint64) (*models.Swipe, error) {
	ret := _m.Called(swiperID, userID)

	var r0 *models.Swipe
	if rf, ok := ret.Get(0).(func(uint64, uint64) *models.Swipe); ok {
		r0 = rf(swiperID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Swipe)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(swiperID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NextUser provides a mock function with given fields: userID
func (_m *Repositories) NextUser(userID uint64) (*models.User, error) {
	ret := _m.Called(userID)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(uint64) *models.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PatchUserLogin provides a mock function with given fields: userID
func (_m *Repositories) PatchUserLogin(userID uint64) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PatchUserLogout provides a mock function with given fields: userID
func (_m *Repositories) PatchUserLogout(userID uint64) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PatchUserVerified provides a mock function with given fields: userID
func (_m *Repositories) PatchUserVerified(userID uint64) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RedisUserGetSwipes provides a mock function with given fields: username
func (_m *Repositories) RedisUserGetSwipes(username string) (int, error) {
	ret := _m.Called(username)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RedisUserRemoveLImit provides a mock function with given fields: username
func (_m *Repositories) RedisUserRemoveLImit(username string) error {
	ret := _m.Called(username)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RedisUserSetSwipes provides a mock function with given fields: username, swipes, t
func (_m *Repositories) RedisUserSetSwipes(username string, swipes int, t time.Duration) error {
	ret := _m.Called(username, swipes, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, time.Duration) error); ok {
		r0 = rf(username, swipes, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepositories interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositories creates a new instance of Repositories. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositories(t mockConstructorTestingTNewRepositories) *Repositories {
	mock := &Repositories{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
