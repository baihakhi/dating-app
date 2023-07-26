package services_test

import (
	"errors"
	"testing"
	"time"

	"github.com/baihakhi/dating-app/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateSwipe(t *testing.T) {
	type expectedData struct {
		swipeLeft, timeLimit int
		result int64
		username string
		errRedisGetSwipe, errRedisSetSwipe, errRedisSetSwipe2, errCreateSwipe error
		lastLogin *time.Time
		swipeData *models.Swipe
	}

	helperTest := func (ex expectedData) *mocks {
		mc := provideMockRepo()
		mc.mockRepository.On("RedisUserGetSwipes", ex.username).
			Return(ex.swipeLeft, ex.errRedisGetSwipe).Once()
		mc.mockRepository.On("RedisUserSetSwipes", ex.username, 10, time.Duration(24)).
			Return(ex.errRedisSetSwipe)
		mc.mockRepository.On("RedisUserSetSwipes", ex.username, ex.swipeLeft-1, time.Duration(ex.timeLimit)).
			Return(ex.errRedisSetSwipe2)
		mc.mockRepository.On("CreateSwipe", ex.swipeData.Swiper, ex.swipeData.Swiped, ex.swipeData.IsLiked).
			Return(ex.result, ex.errCreateSwipe).Once()
		return mc
	}

	tests := []struct {
		name string
		expected expectedData
		funcExpectedService func(t *testing.T, output int64, err error)
	}{
		{
			name: "Should get error get swipe from repository",
			expected: expectedData{
				username: "user1",
				lastLogin: nil,
				swipeLeft: 1,
				swipeData: &models.Swipe{
					Swiper: 1,
					Swiped: 2,
				},
				errRedisGetSwipe: errors.New("Should get error from repositories - redis: nil"),
			},
			funcExpectedService: func(t *testing.T, output int64, err error) {
				assert.Error(t, err, "Should get error from repositories")
			},
		},
		{
			name: "Should get error get swipe from repository - user had login",
			expected: expectedData{
				username: "user1",
				lastLogin: &time.Time{},
				swipeLeft: 11,
				timeLimit: 24,
				swipeData: &models.Swipe{
					Swiper: 1,
					Swiped: 2,
				},
				errRedisGetSwipe: errors.New("Should get error get swipe from repositories"),
				errRedisSetSwipe: errors.New("Should get error set swipe from repositories"),
			},
			funcExpectedService: func(t *testing.T, output int64, err error) {
				assert.Error(t, err, "Should get error from repositories")
			},
		},
		{
			name: "Should get error get swipe from repository - user had login",
			expected: expectedData{
				username: "user1",
				lastLogin: &time.Time{},
				swipeLeft: 10,
				timeLimit: 0,
				swipeData: &models.Swipe{
					Swiper: 1,
					Swiped: 2,
				},
				errRedisGetSwipe: errors.New("Should get error get swipe from repositories"),
				errRedisSetSwipe2: errors.New("Should get error set swipe from repositories"),
			},
			funcExpectedService: func(t *testing.T, output int64, err error) {
				assert.Error(t, err, "Should get error from repositories")
			},
		},
		{
			name: "Should get error create swipe from repository",
			expected: expectedData{
				username: "user1",
				lastLogin: &time.Time{},
				swipeLeft: 10,
				timeLimit: 0,
				swipeData: &models.Swipe{
					Swiper: 1,
					Swiped: 2,
				},
				errCreateSwipe: errors.New("error creating swipe record to db"),
			},
			funcExpectedService: func(t *testing.T, output int64, err error) {
				assert.Error(t, err, "Should get error from repositories")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := helperTest(tt.expected)
			output, err := provideService(mc).CreateSwipe(tt.expected.username, tt.expected.swipeData, tt.expected.lastLogin)
			tt.funcExpectedService(t, output, err)
		})
	}
}

func TestService_GetSwipe(t *testing.T) {
	type expectedData struct {
		swiperID, userID uint64
		err error
		result *models.Swipe
	}

	helperTest := func (ex expectedData) *mocks {
		mc := provideMockRepo()
		mc.mockRepository.On("GetSwipe", ex.swiperID, ex.userID).
			Return(ex.result, ex.err)
		return mc
	}

	tests := []struct{
		name string
		expected expectedData
		funcExpectedService func(t *testing.T, output *models.Swipe, err error)
	}{
		{
			name: "SHould get error get swipe from repositories",
			expected: expectedData{
				swiperID: 1,
				userID: 2,
				err: errors.New("error get swipe"),
			},
			funcExpectedService: func(t *testing.T, output *models.Swipe, err error) {
				assert.Error(t, err, "error get swipe from repositories")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := helperTest(tt.expected)
			output, err := provideService(mc).GetSwipe(tt.expected.swiperID, tt.expected.userID)
			tt.funcExpectedService(t, output, err)
		})
	}
}