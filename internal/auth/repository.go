package auth

import (
	"auth-with-clean-architecture/internal/user"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) FindByUsername(username string) (*user.User, error) {
	var user *user.User
	res := r.db.Where("username = ?", username).First(&user).Error

	return user, res
}

func (r Repository) Login(user *user.User) (*ProfileItemWithToken, error) {
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

func (r Repository) ShowProfile(tokenSigned string) (*user.User, error) {
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
