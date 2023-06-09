package auth

import (
	"auth-with-clean-architecture/internal/user"
	"auth-with-clean-architecture/pkg/password"
	"errors"
)

type UseCase struct {
	repository *Repository
}

func NewUseCase(r *Repository) *UseCase {
	return &UseCase{
		repository: r,
	}
}

func (u UseCase) Login(payload *Payload) (*ProfileItemWithToken, error) {
	user, _ := u.repository.FindByUsername(payload.Username)
	if user.Username == "" {
		return nil, errors.New("user not found")
	}

	match := password.CheckPasswordHash(payload.Password, user.Password)
	if !match {
		return nil, errors.New("password is incorrect")
	}

	return u.repository.Login(user)
}

func (u UseCase) ShowProfile(tokenSigned string) (*user.User, error) {
	return u.repository.ShowProfile(tokenSigned)
}
