package service_test

import (
	"errors"
	"testing"

	"github.com/baihakhi/dating-app/internal/models"
	repository "github.com/baihakhi/dating-app/internal/repositories/mocks"
	service "github.com/baihakhi/dating-app/internal/services"
	"github.com/stretchr/testify/assert"
)

type mocks struct {
	mockRepository repository.Repositories
}

func provideMockRepo() *mocks {
	return &mocks{
		repository.Repositories{},
	}
}

func provideService(m *mocks) service.Services {
	return service.InitService(&m.mockRepository)
}

func TestService_CreateUser(t *testing.T) {
	type expectedData struct {
		username string
		data     *models.User
		err      error
	}

	helperTest := func(ex expectedData) *mocks {
		mc := provideMockRepo()
		mc.mockRepository.On("CreateUser", ex.data).
			Return(ex.username, ex.err).Once()
		return mc
	}

	tests := []struct {
		name                string
		expected            expectedData
		funcExpectedService func(t *testing.T, output string, err error)
	}{
		{
			name: "Should get error repository",
			expected: expectedData{
				username: "user1",
				data: &models.User{
					Username: "user1",
				},
				err: errors.New("error from db"),
			},
			funcExpectedService: func(t *testing.T, output string, err error) {
				assert.Error(t, err, "Should have error")
			},
		},
		{
			name: "Should get succes insert data to db and return username",
			expected: expectedData{
				username: "user1",
				data: &models.User{
					Username: "user1",
				},
				err: nil,
			},
			funcExpectedService: func(t *testing.T, output string, err error) {
				assert.NoError(t, err, "Should no have error")
				assert.NotNil(t, output, "Should username is not nil")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := helperTest(tt.expected)
			output, err := provideService(mc).CreateUser(tt.expected.data)
			tt.funcExpectedService(t, output, err)
		})
	}
}
