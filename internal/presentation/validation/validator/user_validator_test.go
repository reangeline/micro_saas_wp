package validator_test

import (
	"testing"

	"github.com/reangeline/micro_saas/internal/dto"
	"github.com/reangeline/micro_saas/internal/presentation/validation/validator"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnInvalidEmail_WhenValidateUser_ThenShouldReciveAnError(t *testing.T) {
	valid := validator.NewUserValidator()

	user := &dto.UserInput{
		Name:     "abc",
		LastName: "abc",
		Email:    "abc",
	}

	err := valid.ValidateUser(user)

	assert.Error(t, err)

}

func TestGivenAnEmptyName_WhenValidateUser_ThenShouldReciveAnError(t *testing.T) {
	valid := validator.NewUserValidator()

	user := &dto.UserInput{
		Name:     "",
		LastName: "aaaa",
		Email:    "john@doe.com",
	}

	err := valid.ValidateUser(user)

	assert.Error(t, err)

}

func TestGivenAnEmptyLastName_WhenValidateUser_ThenShouldReciveAnError(t *testing.T) {
	valid := validator.NewUserValidator()

	user := &dto.UserInput{
		Name:     "aaaa",
		LastName: "",
		Email:    "john@doe.com",
	}

	err := valid.ValidateUser(user)

	assert.Error(t, err)

}

func TestGivenValidFields_WhenValidateUser_ThenShouldNotReciveAnError(t *testing.T) {
	valid := validator.NewUserValidator()

	user := &dto.UserInput{
		Name:     "aaaa",
		LastName: "xxxx",
		Email:    "john@doe.com",
	}

	err := valid.ValidateUser(user)

	assert.NoError(t, err)

}
