package protocol

import "github.com/reangeline/micro_saas/internal/dto"

type UserValidatorInterface interface {
	ValidateUser(user *dto.UserInput) error
	ValidateUserEmail(email string) error
}
