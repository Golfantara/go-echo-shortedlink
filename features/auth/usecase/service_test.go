package usecase

import (
	"errors"
	"shortlink/features/auth"
	"shortlink/features/auth/dtos"
	"shortlink/features/auth/mocks"
	mockHelpers "shortlink/helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T){
	jwt := mockHelpers.NewJWTInterface(t)
	hash := mockHelpers.NewHashInterface(t)
	repo := mocks.NewRepository(t)
	service := New(repo, jwt, hash)

	var newUsers = auth.Users{
		ID: 0,
		Fullname: "dida",
		PhoneNumber: "4312",
		Email: "dida@example.com",
		Password: "dida123",
	}

	var invalidData = auth.Users{
		Fullname: "dejan",
	}
	t.Run("Succes create", func(t *testing.T){
		hash.On("HashPassword", newUsers.Password).Return("dida123", nil).Once()
		jwtResult := map[string]any{"access_token":"refresh_token"}
		jwt.On("GenerateJWT", mock.Anything).Return(jwtResult).Once()
		newUsers.Password = "dida123"
		repo.On("Insert", &newUsers).Return(&newUsers, nil).Once()

		result, err := service.Create(dtos.InputUsers{
			Fullname: newUsers.Fullname,
			PhoneNumber: newUsers.PhoneNumber,
			Email: newUsers.Email,
			Password: newUsers.Password,
		})

		assert.Nil(t, err)
		assert.Equal(t, newUsers.Fullname, result.Fullname)
		hash.AssertExpectations(t)
		repo.AssertExpectations(t)
	})

	t.Run("validation error", func(t *testing.T) {
		result, err := service.Create(dtos.InputUsers{
			Fullname: invalidData.Fullname,
		})
		assert.Error(t, err)
		assert.EqualError(t, err, "validation failed please check your input and try again")
		assert.Nil(t, result)
		repo.AssertExpectations(t)
	})

	t.Run("Hash password failed", func(t *testing.T) {
		hash.On("HashPassword", newUsers.Password).Return("", errors.New("hash password failed")).Once()

		result, err := service.Create(dtos.InputUsers{
			Fullname: newUsers.Email,
			PhoneNumber: newUsers.PhoneNumber,
			Email: newUsers.Email,
			Password: newUsers.Password,
		})
		assert.Error(t, err)
		assert.EqualError(t, err, "hash password failed")
		assert.Nil(t, result)

		hash.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	jwt := mockHelpers.NewJWTInterface(t)
	hash := mockHelpers.NewHashInterface(t)
	repository := mocks.NewRepository(t)
	service := New(repository, jwt, hash)

	usersData := auth.Users{
		ID: 0,
		Fullname: "dida",
		PhoneNumber: "4312",
		Email: "dida@example.com",
		Password: "dida123",
	}
	loginData := dtos.LoginUsers{
		Email: "dida@example.com",
		Password: "dida123",
	}
	t.Run("succes login", func(t *testing.T){
		jwtResult := map[string]any{"acces_token":"refresh_token"}
		hash.On("CompareHash", usersData.Password,loginData.Password).Return(true).Once()
		jwt.On("GenerateJWT", mock.Anything).Return(jwtResult).Once()
		repository.On("Login", usersData.Email, usersData.Password).Return(&usersData, nil).Once()

		result, err := service.Login(usersData.Email, usersData.Password)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "dida", result.Name)
		assert.Equal(t, jwtResult, result.Token)
		repository.AssertExpectations(t)
		hash.AssertExpectations(t)
		jwt.AssertExpectations(t)
	})
}

func TestFindAll(t *testing.T) {
	jwt := mockHelpers.NewJWTInterface(t)
	hash := mockHelpers.NewHashInterface(t)
	repository := mocks.NewRepository(t)
	service := New(repository, jwt, hash)

	userEntities := []auth.Users{
		{ID:1, Fullname: "Users 1"},
		{ID:2, Fullname: "Users 2"},
	}
	repository.On("Paginate", 1, 10).Return(userEntities).Once()
	users := service.FindAll(1, 10)

	expected := []dtos.ResUsers{
		{ID: 1, Fullname: "Users 1"},
		{ID: 2, Fullname: "Users 2"},
	}

	assert.Equal(t, expected, users)
	repository.AssertExpectations(t)
}

func TestRemove(t *testing.T){
	jwt := mockHelpers.NewJWTInterface(t)
	hash := mockHelpers.NewHashInterface(t)
	repository := mocks.NewRepository(t)
	service := New(repository, jwt, hash)

	userID := 1
	repository.On("DeleteByID", userID).Return(int64(1)).Once()
	result := service.Remove(userID)

	assert.True(t, result)

	repository.AssertExpectations(t)
}

func TestModify(t *testing.T){
	jwt := mockHelpers.NewJWTInterface(t)
	hash := mockHelpers.NewHashInterface(t)
	repository := mocks.NewRepository(t)
	service := New(repository, jwt, hash)

	var updateData = dtos.InputUsers{
		Fullname: "dida",
		PhoneNumber: "324",
		Email: "didac@example.com",
		Password: "diddda",
	}

	var users = auth.Users{
		ID: 1,
		Fullname: "dejan",
		PhoneNumber: "324334",
		Email: "dejan@example.com",
		Password: "dejan11",
	}
	t.Run("succes update data", func(t *testing.T){
		repository.On("Update", mock.Anything).Return(int64(1)).Once()

		result := service.Modify(updateData, users.ID)
		assert.True(t, result)
		repository.AssertExpectations(t)
	})
}


func TestFindByID(t *testing.T) {
	jwt := mockHelpers.NewJWTInterface(t)
	hash := mockHelpers.NewHashInterface(t)
	repository := mocks.NewRepository(t)
	service := New(repository, jwt, hash)

	userID := 1

	userEntity := &auth.Users{
		ID: 1,
		Fullname: "User 1",
	}

	t.Run("success find by ID", func(t *testing.T) {
		repository.On("SelectByID", userID).Return(userEntity).Once()

		result := service.FindByID(userID)

		expected := &dtos.ResUsers{
			ID: 1,
			Fullname: "User 1",
		}

		assert.Equal(t, expected, result)

		repository.AssertExpectations(t)
	})
}
