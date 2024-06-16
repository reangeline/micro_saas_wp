package usecase

import (
	"context"
	"errors"

	"github.com/reangeline/micro_saas/internal/domain/contract/repository"
	"github.com/reangeline/micro_saas/internal/domain/entity"
	"github.com/reangeline/micro_saas/internal/dto"
)

type UserUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserUseCase(
	userRepository repository.UserRepositoryInterface,
) *UserUseCase {
	return &UserUseCase{
		userRepository,
	}
}

var (
	ErrEmailAlreadyExists = errors.New("email already exist")
)

func (u *UserUseCase) CreateUser(ctx context.Context, input *dto.UserInput) error {

	user, err := entity.NewUser(input.Name, input.LastName, input.Email)
	if err != nil {
		return err
	}

	userModel := &entity.User{
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}

	if err := u.userRepository.CreateUser(ctx, userModel); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) FindUserByEmail(email string) (*dto.UserOutput, error) {
	user, err := u.userRepository.FindByUserEmail(email)

	if err != nil {
		return nil, err
	}

	return &dto.UserOutput{
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}, nil

}

func (u *UserUseCase) FindAll() ([]*dto.UserOutput, error) {

	users, err := u.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var userOutput []*dto.UserOutput
	for _, value := range users {
		userOutput = append(userOutput, &dto.UserOutput{
			Name:     value.Name,
			LastName: value.LastName,
			Email:    value.Email,
		})
	}

	return userOutput, nil
}

func (u *UserUseCase) UpdateByEmail(input *dto.UserInput) (*dto.UserOutput, error) {

	user, err := entity.NewUser(input.Name, input.LastName, input.Email)

	if err != nil {
		return nil, err
	}

	userEntity, err := u.userRepository.UpdateByEmail(user)
	if err != nil {
		return nil, err
	}

	return &dto.UserOutput{
		Name:     userEntity.Name,
		LastName: userEntity.LastName,
		Email:    userEntity.Email,
	}, nil
}
