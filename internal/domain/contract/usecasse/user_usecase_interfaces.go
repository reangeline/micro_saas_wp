package usecase

import (
	"context"

	"github.com/reangeline/micro_saas/internal/dto"
)

type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, input *dto.UserInput) error
	FindAll() ([]*dto.UserOutput, error)
	FindUserByEmail(email string) (*dto.UserOutput, error)
	UpdateByEmail(input *dto.UserInput) (*dto.UserOutput, error)
}
