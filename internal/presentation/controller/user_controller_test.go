package controller_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	usecase "github.com/reangeline/micro_saas/internal/domain/contract/usecasse"
	"github.com/reangeline/micro_saas/internal/dto"
	"github.com/reangeline/micro_saas/internal/presentation/controller"
	"github.com/stretchr/testify/assert"
)

// MockUserUseCase is a mock implementation of usecase.UserUseCaseInterface
type MockUserUseCase struct {
	usecase.UserUseCaseInterface
}

func (m *MockUserUseCase) CreateUser(ctx context.Context, user *dto.UserInput) error {
	return nil
}

// MockUserValidator is a mock implementation of protocols.UserValidatorInterface
type MockUserValidator struct {
}

func (m *MockUserValidator) ValidateUser(user *dto.UserInput) error {
	return nil
}

func (m *MockUserValidator) ValidateUserEmail(email string) error {
	return nil
}

func TestCreateUserController_CreateUser(t *testing.T) {

	userUseCase := &MockUserUseCase{}
	// userValidator := &MockUserValidator{}

	userController := controller.NewUserController(userUseCase)

	t.Run("Successful user creation", func(t *testing.T) {
		requestBody := `{"name": "John", "email": "john@example.com"}`

		req, err := http.NewRequest("POST", "/users", bytes.NewReader([]byte(requestBody)))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		userController.CreateUserRest(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("Invalid request body", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/users", bytes.NewReader([]byte("invalid-json")))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		userController.CreateUserRest(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}
