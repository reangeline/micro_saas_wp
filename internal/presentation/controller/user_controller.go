package controller

import (
	"context"
	"encoding/json"
	"net/http"

	usecase "github.com/reangeline/micro_saas/internal/domain/contract/usecasse"
	"github.com/reangeline/micro_saas/internal/dto"
)

type Error struct {
	Message string `json:"message"`
}

type UserController struct {
	userUseCase usecase.UserUseCaseInterface
}

func NewUserController(
	userUseCase usecase.UserUseCaseInterface,
) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.UserInput  true  "user request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /users [post]
func (u *UserController) CreateUserRest(w http.ResponseWriter, r *http.Request) {
	var user dto.UserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	ctx := r.Context()
	err = u.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (u *UserController) CreateUser(ctx context.Context, input dto.UserInput) error {

	// err := u.userValidator.ValidateUser(&input)

	// if err != nil {
	// 	return err
	// }

	err := u.userUseCase.CreateUser(ctx, &input)

	if err != nil {
		return err
	}

	return nil
}

// @Summary      Find user by email
// @Description  Find user by email
// @Tags         find users
// @Accept       json
// @Produce      json
// @Param        email query string true "Endereço de email do usuário"
// @Success      200 {object} dto.UserOutput
// @Failure      400 {object} Error
// @Router       /users [get]
func (u *UserController) FindUserByEmailRest(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("error")
	}

	user, err := u.userUseCase.FindUserByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

// @Summary      Find user by email
// @Description  Find user by email
// @Tags         find users
// @Accept       json
// @Produce      json
// @Param        email query string true "Endereço de email do usuário"
// @Success      200 {object} dto.UserOutput
// @Failure      400 {object} Error
// @Router       /users [get]
func (u *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	// email := r.URL.Query().Get("email")

	// ctx := r.Context()

	user, err := u.userUseCase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (u *UserController) UpdateByEmail(w http.ResponseWriter, r *http.Request) {
	var user dto.UserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	userOut, err := u.userUseCase.UpdateByEmail(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userOut)
}
