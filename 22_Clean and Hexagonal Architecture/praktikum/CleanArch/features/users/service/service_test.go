package service

import (
	"errors"
	"restEcho1/features/users"
	"restEcho1/features/users/mocks"
	helper "restEcho1/helper/mocks"

	"github.com/stretchr/testify/mock"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	generator := helper.NewGeneratorInterface(t)
	jwt := helper.NewJWTInterface(t)
	data := mocks.NewUserDataInterface(t)
	service := New(data, generator, jwt)
	newUser := users.User{
		Nama:     "jerry",
		HP:       "12345",
		Password: "mantul123",
	}

	t.Run("Success insert", func(t *testing.T) {
		generator.On("GenerateUUID").Return("randomUUID", nil).Once()
		newUser.ID = "randomUUID"
		data.On("Insert", newUser).Return(&newUser, nil).Once()

		result, err := service.Register(newUser)
		assert.Nil(t, err)
		assert.Equal(t, newUser.ID, result.ID)
		assert.Equal(t, newUser.Nama, result.Nama)
		generator.AssertExpectations(t)
		data.AssertExpectations(t)
	})

	t.Run("Generate failed", func(t *testing.T) {
		generator.On("GenerateUUID").Return("", errors.New("some error on generator")).Once()

		result, err := service.Register(newUser)
		assert.Error(t, err)
		assert.EqualError(t, err, "id generator failed")
		assert.Nil(t, result)
		generator.AssertExpectations(t)
	})

	// t.Run("Insert failed", func(t *testing.T) {
	// Coba dilanjutkan ya
	// })
}

func TestLogin(t *testing.T) {
	generator := helper.NewGeneratorInterface(t)
	j := helper.NewJWTInterface(t)
	data := mocks.NewUserDataInterface(t)
	service := New(data, generator, j)
	userData := users.User{
		ID:       "randomUserID",
		Nama:     "jerry",
		HP:       "12345",
		Password: "mantul123",
	}

	t.Run("success login", func(t *testing.T) {
		jwtResult := map[string]any{"access_token": "randomAccessToken"}
		data.On("Login", userData.HP, userData.Password).Return(&userData, nil)
		j.On("GenerateJWT", mock.Anything).Return(jwtResult)
		result, err := service.Login(userData.HP, userData.Password)

		data.AssertExpectations(t)
		j.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "jerry", result.Nama)
		assert.Equal(t, jwtResult, result.Access)
	})
}

func TestGetAllUsers(t *testing.T) {
	generator := helper.NewGeneratorInterface(t)
	j := helper.NewJWTInterface(t)
	data := mocks.NewUserDataInterface(t)
	service := New(data, generator, j)

	t.Run("success get all users", func(t *testing.T) {
		// Define the mock behavior for GetAllUsers
		expectedUsers := []users.User{
			{
				ID:       "1",
				Nama:     "User1",
				HP:       "12345",
				Password: "pass1",
			},
			{
				ID:       "2",
				Nama:     "User2",
				HP:       "67890",
				Password: "pass2",
			},
			// Add more mock user data as needed
		}

		data.On("GetAllUsers").Return(expectedUsers, nil)

		// Call the GetAllUsers method
		result, err := service.GetAllUsers()

		data.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedUsers, result)
	})

	t.Run("error getting users", func(t *testing.T) {
		// Define the mock behavior for GetAllUsers when there's an error
		data.On("GetAllUsers").Return(nil, errors.New("database error"))

		// Call the GetAllUsers method
		result, err := service.GetAllUsers()

		data.AssertExpectations(t)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
