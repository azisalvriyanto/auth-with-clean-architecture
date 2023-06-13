package customer

import (
	"auth-with-clean-architecture/internal/customer"
	"auth-with-clean-architecture/internal/customer/mocks"
	"errors"
	"reflect"
	"testing"
)

var (
	errorCase = errors.New("some error")
	existData = customer.Customer{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Avatar:    "",
	}
	nilData = customer.Customer{}
)

func TestUseCaseShowAll(t *testing.T) {
	type fields struct {
		R customer.RepositoryInterface
	}

	successCase := []customer.Customer{}

	for i := 0; i < 5; i++ {
		successCase = append(successCase, existData)
	}

	mockRepository := mocks.NewRepositoryInterface(t)
	mockRepository.EXPECT().ShowAll().Return(successCase, nil).Once()
	mockRepository.EXPECT().ShowAll().Return(nil, errorCase).Once()

	tests := []struct {
		name    string
		fields  fields
		want    []customer.Customer
		wantErr bool
	}{
		{
			name: "success case",
			fields: fields{
				R: mockRepository,
			},
			want:    successCase,
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

func TestUseCaseCreate(t *testing.T) {
	type fields struct {
		R customer.RepositoryInterface
	}
	type args struct {
		customer *customer.Customer
	}

	mockRepository := mocks.NewRepositoryInterface(t)
	mockRepository.EXPECT().Create(&existData).Return(nil).Once()
	mockRepository.EXPECT().Create(&nilData).Return(errorCase).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				customer: &existData,
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				customer: &nilData,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &customer.UseCase{
				R: tt.fields.R,
			}
			if err := u.Create(tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCaseShow(t *testing.T) {
	type fields struct {
		R customer.RepositoryInterface
	}
	type args struct {
		ID string
	}

	mockRepository := mocks.NewRepositoryInterface(t)
	mockRepository.EXPECT().Show("1").Return(&existData, nil).Once()
	mockRepository.EXPECT().Show("").Return(&nilData, errorCase).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customer.Customer
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				ID: "1",
			},
			want:    &existData,
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				ID: "",
			},
			want:    &nilData,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &customer.UseCase{
				R: tt.fields.R,
			}
			got, err := u.Show(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Show() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.Show() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseUpdate(t *testing.T) {
	type fields struct {
		R customer.RepositoryInterface
	}
	type args struct {
		ID       string
		customer customer.Customer
	}

	mockRepository := mocks.NewRepositoryInterface(t)
	mockRepository.EXPECT().Update("1", existData).Return(&existData, nil).Once()
	mockRepository.EXPECT().Update("1", nilData).Return(&nilData, errorCase).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customer.Customer
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				ID:       "1",
				customer: existData,
			},
			want:    &existData,
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				ID:       "1",
				customer: nilData,
			},
			want:    &nilData,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &customer.UseCase{
				R: tt.fields.R,
			}
			got, err := u.Update(tt.args.ID, tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseDestroy(t *testing.T) {
	type fields struct {
		R customer.RepositoryInterface
	}
	type args struct {
		ID string
	}

	mockRepository := mocks.NewRepositoryInterface(t)
	mockRepository.EXPECT().Destroy("1").Return(&existData, nil).Once()
	mockRepository.EXPECT().Destroy("").Return(&nilData, errorCase).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customer.Customer
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				ID: "1",
			},
			want:    &existData,
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				ID: "",
			},
			want:    &nilData,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &customer.UseCase{
				R: tt.fields.R,
			}
			got, err := u.Destroy(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Destroy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.Destroy() = %v, want %v", got, tt.want)
			}
		})
	}
}
