package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Controller struct {
	useCase *UseCase
}

func NewController(useCase *UseCase) *Controller {
	return &Controller{
		useCase: useCase,
	}
}

type UserItem struct {
	ID       uint   `json:"id"`
	FullName string `json:"first_name"`
	Username string `json:"last_name"`
	Password string `json:"password"`
}

type UserItemResponse struct {
	Message string    `json:"message"`
	Data    *UserItem `json:"data"`
}

func (c Controller) ShowAll() (*[]UserItem, error) {
	users, err := c.useCase.ShowAll()
	if err != nil {
		return nil, err
	}

	res := &[]UserItem{}
	for _, user := range users {
		c := UserItem{
			ID:       user.ID,
			FullName: user.FullName,
			Username: user.Username,
			Password: user.Password,
		}
		*res = append(*res, c)
	}

	return res, nil
}

func (c Controller) Create(body *CreateRequest) (*UserItemResponse, error) {
	user := User{
		Model:    gorm.Model{},
		FullName: body.FullName,
		Username: body.Username,
		Password: body.Password,
	}
	err := c.useCase.Create(&user)
	if err != nil {
		return nil, err
	}

	res := &UserItemResponse{
		Message: "user successfully created",
		Data: &UserItem{
			ID:       user.ID,
			FullName: body.FullName,
			Username: body.Username,
			Password: body.Password,
		},
	}

	return res, nil
}
func (c Controller) Show(ID string) (*UserItem, error) {
	user, err := c.useCase.Show(ID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	res := &UserItem{
		ID:       user.ID,
		FullName: user.FullName,
		Username: user.Username,
		Password: user.Password,
	}

	return res, nil
}

func (c Controller) Update(ID string, body User) (*UserItemResponse, error) {
	user, err := c.useCase.Update(ID, body)
	if err != nil {
		return nil, err
	}

	res := &UserItemResponse{
		Message: "user successfully updated",
		Data: &UserItem{
			ID:       user.ID,
			FullName: user.FullName,
			Username: user.Username,
			Password: user.Password,
		},
	}

	return res, nil
}

func (c Controller) Destroy(ID string) (*UserItemResponse, error) {
	user, err := c.useCase.Destroy(ID)
	if err != nil {
		return nil, err
	}

	res := &UserItemResponse{
		Message: "user successfully destroyed",
		Data: &UserItem{
			ID:       user.ID,
			FullName: user.FullName,
			Username: user.Username,
			Password: user.Password,
		},
	}

	return res, nil
}
