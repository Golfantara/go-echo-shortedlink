package usecase

import (
	"errors"
	"shortlink/features/donate"
	"shortlink/features/donate/dtos"
	"shortlink/features/donate/mocks"
	mockHelpers "shortlink/helpers/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	repo := mocks.NewRepository(t)
	validate := validator.New()
	generator := mockHelpers.NewGeneratorInterface(t)
	service := New(repo, validate, generator)

	t.Run("succes insert", func(t *testing.T){
		newData := []donate.Transaction{
			{ID: 1, UserID: "1", OrderID: "randomUUID"},
			{ID: 2, UserID: "2", OrderID: "randomUUID"},
		}
		repo.On("Paginate", 1, 10).Return(newData).Once()
		transactions := service.FindAll(1, 10)
	
		expected := []dtos.TransactionInputResponse{
			{ID: 1, UserID: "1", OrderID: "randomUUID"},
			{ID: 2, UserID: "2", OrderID: "randomUUID"},
		}
	
		assert.Equal(t, expected, transactions)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T){
	repo := mocks.NewRepository(t)
	validate := validator.New()
	generator := mockHelpers.NewGeneratorInterface(t)
	service := New(repo, validate, generator)

	newData := donate.Transaction{
		UserID: "3",
		Amount: 10000,
		Description: "ini untuk jajan",
	}
	t.Run("succes insert", func(t *testing.T){
		newData.OrderID = "randomUUID"

		generator.On("GenerateUUID").Return("randomUUID", nil).Once()
		repo.On("Insert", &newData).Return(&newData, nil).Once()
		repo.On("SnapRequest", "randomUUID", int64(10000)).Return("randomToken", "https://random.com").Once()


		result, err := service.Create(dtos.TransactionInput{
			UserID: newData.UserID,
			Amount: uint(newData.Amount),
			Description: newData.Description,
		})
		assert.Nil(t, err)
		assert.Equal(t, "randomUUID", result.OrderID)
		generator.AssertExpectations(t)
		repo.AssertExpectations(t)

	})
}

func TestNotification(t *testing.T){
	repo := mocks.NewRepository(t)
	validate := validator.New()
	generator := mockHelpers.NewGeneratorInterface(t)
	service := New(repo, validate, generator)

	orderID := "randomOrderID"
	notificationPayload := map[string]any{
		"order_id":     orderID,
	}

	expectedErr := errors.New("Check transactions error")
	repo.On("CheckTransaction",orderID).Return(donate.Status{}, expectedErr).Once()

	err := service.Notifications(notificationPayload)
	assert.Equal(t, expectedErr, err)
	repo.AssertExpectations(t)
}