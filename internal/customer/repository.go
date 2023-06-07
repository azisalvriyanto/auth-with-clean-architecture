package customer

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) ShowAll() ([]Customer, error) {
	var customers []Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r Repository) Create(customer *Customer) error {
	return r.db.Create(customer).Error
}

func (r Repository) Show(ID string) (*Customer, error) {
	var customer Customer
	err := r.db.First(&customer, ID).Error
	return &customer, err
}

func (r Repository) Update(ID string, body Customer) (*Customer, error) {
	var customer Customer
	err := r.db.First(&customer, ID).Error
	customer.FirstName = body.FirstName
	customer.LastName = body.LastName
	customer.Email = body.Email
	customer.Avatar = body.Avatar

	query := r.db.Save(&customer).Error
	if query != nil {
		return nil, query
	}

	return &customer, err
}

func (r Repository) Destroy(ID string) (*Customer, error) {
	var customer Customer
	err := r.db.First(&customer, ID).Error

	query := r.db.Delete(&customer).Error
	if query != nil {
		return nil, query
	}

	return &customer, err
}
