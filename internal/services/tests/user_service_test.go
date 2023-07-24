package services_test

import (
	"errors"
	"testing"
	"time"

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
			name: "Should get error from repository",
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

func TestService_GetOneUserByUsername(t *testing.T) {
	type expectedData struct {
		username string
		data     *models.User
		err      error
	}

	helperTest := func(ex expectedData) *mocks {
		mc := provideMockRepo()
		mc.mockRepository.On("GetOneUsersByUsername", ex.username).
			Return(ex.data, ex.err).Once()
		return mc
	}

	tests := []struct {
		name                string
		expected            expectedData
		funcExpectedService func(t *testing.T, output *models.User, err error)
	}{
		{
			name: "Should get error from repository",
			expected: expectedData{
				username: "user1",
				data: &models.User{
					Username: "user1",
				},
				err: errors.New("error from db"),
			},
			funcExpectedService: func(t *testing.T, output *models.User, err error) {
				assert.Error(t, err, "Should have error")
			},
		},
		{
			name: "Should get succes get data from db and return user",
			expected: expectedData{
				username: "user1",
				data: &models.User{},
				err: nil,
			},
			funcExpectedService: func(t *testing.T, output *models.User, err error) {
				assert.NoError(t, err, "Should no have error")
				assert.NotNil(t, output, "Should username is not nil")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := helperTest(tt.expected)
			output, err := provideService(mc).GetOneUserByUsername(tt.expected.username)
			tt.funcExpectedService(t, output, err)
		})
	}
}

func TestService_PatchUserVerified(t *testing.T) {
	type expectedData struct {
		userID uint64
		err      error
	}

	helperTest := func(ex expectedData) *mocks {
		mc := provideMockRepo()
		mc.mockRepository.On("PatchUserVerified", ex.userID).
			Return(ex.err).Once()
		return mc
	}

	tests := []struct {
		name                string
		expected            expectedData
		funcExpectedService func(t *testing.T, err error)
	}{
		{
			name: "Should get error from repository",
			expected: expectedData{
				userID: 1,
				err: errors.New("error from db"),
			},
			funcExpectedService: func(t *testing.T, err error) {
				assert.Error(t, err, "Should have error")
			},
		},
		{
			name: "Should get succes update data to db",
			expected: expectedData{
				userID: 1,
				err: nil,
			},
			funcExpectedService: func(t *testing.T, err error) {
				assert.NoError(t, err, "Should no have error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := helperTest(tt.expected)
			err := provideService(mc).PatchUserVerified(tt.expected.userID)
			tt.funcExpectedService(t, err)
		})
	}
}

func TestService_RemoveSwipeLimit(t *testing.T) {
	type expectedData struct {
		username string
		err      error
	}

	helperTest := func(ex expectedData) *mocks {
		mc := provideMockRepo()
		mc.mockRepository.On("RedisUserRemoveLImit", ex.username).
			Return(ex.err).Once()
		return mc
	}

	tests := []struct {
		name                string
		expected            expectedData
		funcExpectedService func(t *testing.T, err error)
	}{
		{
			name: "Should get error from repository",
			expected: expectedData{
				username: "user1",
				err: errors.New("error from db"),
			},
			funcExpectedService: func(t *testing.T, err error) {
				assert.Error(t, err, "Should have error")
			},
		},
		{
			name: "Should get succes update data to db",
			expected: expectedData{
				username: "user1",
				err: nil,
			},
			funcExpectedService: func(t *testing.T, err error) {
				assert.NoError(t, err, "Should no have error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := helperTest(tt.expected)
			err := provideService(mc).RemoveSwipeLimit(tt.expected.username)
			tt.funcExpectedService(t, err)
		})
	}
}

func TestService_NextUser(t *testing.T) {
	type expectedData struct {
		userID uint64
		result *models.User
		err      error
	}

	helperTest := func(ex expectedData) *mocks {
		mc := provideMockRepo()
		mc.mockRepository.On("NextUser", ex.userID).
			Return(ex.result, ex.err).Once()
		return mc
	}

	tests := []struct {
		name                string
		expected            expectedData
		funcExpectedService func(t *testing.T, output *models.User, err error)
	}{
		{
			name: "Should get error from repository",
			expected: expectedData{
				userID: 1,
				err: errors.New("error from db"),
			},
			funcExpectedService: func(t *testing.T, output *models.User, err error) {
				assert.Error(t, err, "Should have error")
			},
		},
		{
			name: "Should get succes get data from db and return user",
			expected: expectedData{
				userID: 1,
				result: &models.User{},
				err: nil,
			},
			funcExpectedService: func(t *testing.T, output *models.User, err error) {
				assert.NoError(t, err, "Should no have error")
				assert.NotNil(t, output, "Should user is not nil")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := helperTest(tt.expected)
			res, err := provideService(mc).NextUser(tt.expected.userID)
			tt.funcExpectedService(t, res, err)
		})
	}
}

func TestService_Login(t *testing.T) {
	type expectedData struct {
		password, result string
		errGetPassword, errGetUser, errPatchUser, errSetRedis      error
		data, userAcc *models.User
	}

	helperTest := func(ex expectedData) *mocks {
		mc := provideMockRepo()
		mc.mockRepository.On("GetPasswordByUsername", ex.data.Username).
			Return(ex.password, ex.errGetPassword).Once()
		mc.mockRepository.On("GetOneUsersByUsername", ex.data.Username).
			Return(ex.userAcc, ex.errGetUser).Once()
		mc.mockRepository.On("PatchUserLogin", ex.data.ID).
			Return(ex.errPatchUser).Once()
		mc.mockRepository.On("RedisUserSetSwipes", ex.data.Username, 10, time.Duration(24)).
			Return(ex.errSetRedis)
		return mc
	}

	tests := []struct {
		name                string
		expected            expectedData
		funcExpectedService func(t *testing.T, output string, err error)
	}{
		{
			name: "Should get error get password from repository",
			expected: expectedData{
				data: &models.User{
					Username: "user1",
				},
				errGetPassword: errors.New("error get password from db"),
			},
			funcExpectedService: func(t *testing.T, output string, err error) {
				assert.Error(t, err, "Should have error")
			},
		},
		{
			name: "Should get error get user from repository",
			expected: expectedData{
				data: &models.User{
					Username: "user1",
					Password: "password",
				},
				password: "$2a$12$eUchbF3KiBg9OOhPYvGg8e1Mkdl0bFMG6VOQl5rwmSWmjSyeAsd4C",
				userAcc: &models.User{
					ID: 1,
					Username: "user1",
				},
				errGetUser: errors.New("error get user from db"),
			},
			funcExpectedService: func(t *testing.T, output string, err error) {
				assert.Error(t, err, "Should have error")
			},
		},
		{
			name: "Should get error patch user from repository",
			expected: expectedData{
				data: &models.User{
					ID: 1,
					Username: "user1",
					Password: "password",
				},
				password: "$2a$12$eUchbF3KiBg9OOhPYvGg8e1Mkdl0bFMG6VOQl5rwmSWmjSyeAsd4C",
				userAcc: &models.User{
					ID: 1,
					Username: "user1",
				},
				errPatchUser: errors.New("error get patch user from db"),
			},
			funcExpectedService: func(t *testing.T, output string, err error) {
				assert.Error(t, err, "Should have error")
			},
		},
		{
			name: "Should get error set swipes from repository",
			expected: expectedData{
				data: &models.User{
					ID: 1,
					Username: "user1",
					Password: "password",
				},
				password: "$2a$12$eUchbF3KiBg9OOhPYvGg8e1Mkdl0bFMG6VOQl5rwmSWmjSyeAsd4C",
				userAcc: &models.User{
					ID: 1,
					Username: "user1",
					IsVerified: false,
				},
				errSetRedis: errors.New("error get set swipe from db"),
			},
			funcExpectedService: func(t *testing.T, output string, err error) {
				assert.Error(t, err, "Should have error")
			},
		},
		{
			name: "Should get succes get data from db and return user",
			expected: expectedData{
				data: &models.User{
					ID: 1,
					Username: "user1",
					Password: "password",
				},
				password: "$2a$12$eUchbF3KiBg9OOhPYvGg8e1Mkdl0bFMG6VOQl5rwmSWmjSyeAsd4C",
				userAcc: &models.User{
					ID: 1,
					Username: "user1",
					IsVerified: false,
				},
				result: "user.token.ex.12387912873hdsjfdsfskj",
			},
			funcExpectedService: func(t *testing.T, output string, err error) {
				assert.NoError(t, err, "Should no have error")
				assert.NotNil(t, output, "Should user is not nil")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := helperTest(tt.expected)
			res, err := provideService(mc).Login(tt.expected.data)
			tt.funcExpectedService(t, res, err)
		})
	}
}

func TestService_Logout(t *testing.T) {
	type expectedData struct {
		errGetUser, errPatchUser error
		data, userAcc *models.User
	}

	helperTest := func(ex expectedData) *mocks {
		mc := provideMockRepo()
		mc.mockRepository.On("GetOneUsersByUsername", ex.data.Username).
			Return(ex.userAcc, ex.errGetUser).Once()
		mc.mockRepository.On("PatchUserLogout", ex.userAcc.ID).
			Return(ex.errPatchUser).Once()
		return mc
	}
	
	tests := []struct {
		name string
		expected expectedData
		funcExpectedService func(t *testing.T, err error)
	}{
		{
			name: "Should get error get user from repositories",
			expected: expectedData{
				data: &models.User{
					Username: "user1",
				},
				userAcc: &models.User{},
				errGetUser: errors.New("error get user from db"),
			},
			funcExpectedService: func(t *testing.T, err error) {
				assert.Error(t, err, "Should get error")
			},
		},
		{
			name: "Should get error patch user from repositories",
			expected: expectedData{
				data: &models.User{
					Username: "user1",
				},
				userAcc: &models.User{
					ID: 1,
					Username: "user1",
				},
				errPatchUser: errors.New("error patch user from db"),
			},
			funcExpectedService: func(t *testing.T, err error) {
				assert.Error(t, err, "Should get error")
			},
		},
		{
			name: "Should success logout user",
			expected: expectedData{
				data: &models.User{
					Username: "user1",
				},
				userAcc: &models.User{
					ID: 1,
					Username: "user1",
				},
			},
			funcExpectedService: func(t *testing.T, err error) {
				assert.NoError(t, err, "Should get no error")
			},
		},
	}

	for _, tt:= range tests{
		t.Run(tt.name, func(t *testing.T) {
			mc := helperTest(tt.expected)
			err := provideService(mc).Logout(tt.expected.data)
			tt.funcExpectedService(t, err)
		})
	}

}