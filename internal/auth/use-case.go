package auth

import (
	"auth-with-clean-architecture/internal/user"
)

type UseCase struct {
	R RepositoryInterface
}
type UseCaseInterface interface {
	Login(payload *Payload) (*ProfileItemWithToken, error)
	ShowProfile(tokenSigned string) (*user.User, error)
}

func NewUseCase(r RepositoryInterface) UseCaseInterface {
	return &UseCase{
		R: r,
	}
}

func (u *UseCase) Login(payload *Payload) (*ProfileItemWithToken, error) {
	return u.R.Login(payload)
}

func (u *UseCase) ShowProfile(tokenSigned string) (*user.User, error) {
	return u.R.ShowProfile(tokenSigned)
}
