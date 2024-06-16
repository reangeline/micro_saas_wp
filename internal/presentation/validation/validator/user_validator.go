package validator

import (
	"strings"

	"github.com/reangeline/micro_saas/internal/dto"
	"github.com/reangeline/micro_saas/internal/presentation/erro"
)

type UserValidator struct {
}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (uv *UserValidator) ValidateUser(user *dto.UserInput) error {

	err := uv.ValidateUserEmail(user.Email)
	if err != nil {
		return err
	}

	if user.Name == "" {
		return erro.ErrNameIsRequired
	}

	if user.LastName == "" {
		return erro.ErrLastNameIsRequired
	}

	return nil
}

func (uv *UserValidator) ValidateUserEmail(email string) error {

	if !strings.Contains(email, "@") {
		return erro.ErrValidEmail
	}

	return nil
}
