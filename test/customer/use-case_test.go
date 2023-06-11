package customer

import (
	"auth-with-clean-architecture/internal/customer"
	"auth-with-clean-architecture/internal/customer/mocks"
	"errors"
	"reflect"
	"testing"
)

func TestUseCase_ShowAll(t *testing.T) {
	type fields struct {
		R customer.RepositoryInterface
	}

	customers := []customer.Customer{}
	errCase := errors.New("some error")

	for i := 0; i < 5; i++ {
		customers = append(customers, customer.Customer{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@example.com",
			Avatar:    "",
		})
	}

	mockRepository := mocks.NewRepositoryInterface(t)
	mockRepository.EXPECT().ShowAll().Return(customers, nil).Once()
	mockRepository.EXPECT().ShowAll().Return(nil, errCase).Once()

	tests := []struct {
		name    string
		fields  fields
		want    []customer.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success case",
			fields: fields{
				R: mockRepository,
			},
			want:    customers,
			wantErr: false,
		},
		{
			name: "error case",
			fields: fields{
				R: mockRepository,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &customer.UseCase{
				R: tt.fields.R,
			}
			got, err := u.ShowAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.ShowAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.ShowAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
