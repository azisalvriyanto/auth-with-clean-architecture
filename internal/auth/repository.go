package auth

import (
	"auth-with-clean-architecture/internal/user"
	"auth-with-clean-architecture/pkg/password"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}
type RepositoryInterface interface {
	FindByUsername(username string) (*user.User, error)
	Login(user *Payload) (*ProfileItemWithToken, error)
	ShowProfile(tokenSigned string) (*user.User, error)
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) FindByUsername(username string) (*user.User, error) {
	var user *user.User
	res := r.DB.Where("username = ?", username).First(&user).Error

	return user, res
}

func (r *Repository) Login(payload *Payload) (*ProfileItemWithToken, error) {
	user, _ := r.FindByUsername(payload.Username)
	if user.Username == "" {
		return nil, errors.New("user not found")
	}

	match := password.CheckPasswordHash(payload.Password, user.Password)
	if !match {
		return nil, errors.New("password is incorrect")
	}

	expTime := time.Now().Add(time.Minute * 3600)
	claims := &JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt-token",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString(JWT_KEY)
	if err != nil {
		return nil, err
	}

	return &ProfileItemWithToken{
		ProfileItem: ProfileItem{
			ID:       user.ID,
			FullName: user.FullName,
			Username: user.Username,
		},
		Token: tokenSigned,
	}, err
}

func (r *Repository) ShowProfile(tokenSigned string) (*user.User, error) {
	token, err := jwt.Parse(tokenSigned, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(JWT_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	res, err := r.FindByUsername(claims["Username"].(string))

	return res, err
}
