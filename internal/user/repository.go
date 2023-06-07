package user

import (
	"auth-with-clean-architecture/pkg/password"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) ShowAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r Repository) Create(user *User) error {
	hash, _ := password.HashPassword(user.Password)
	user.Password = string(hash)
	user.RoleID = 2

	return r.db.Create(user).Error
}

func (r Repository) Show(ID string) (*User, error) {
	var user User
	err := r.db.First(&user, ID).Error
	return &user, err
}

func (r Repository) Update(ID string, body User) (*User, error) {
	var user User
	err := r.db.First(&user, ID).Error
	user.FullName = body.FullName
	user.Username = body.Username
	if user.Password != "" {
		hash, _ := password.HashPassword(body.Password)
		user.Password = string(hash)
	}

	query := r.db.Save(&user).Error
	if query != nil {
		return nil, query
	}

	return &user, err
}

func (r Repository) Destroy(ID string) (*User, error) {
	var user User
	err := r.db.First(&user, ID).Error

	query := r.db.Delete(&user).Error
	if query != nil {
		return nil, query
	}

	return &user, err
}
