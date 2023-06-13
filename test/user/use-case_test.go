package user

import (
	"auth-with-clean-architecture/internal/user"
	"auth-with-clean-architecture/internal/user/mocks"
	"errors"
	"reflect"
	"testing"
)

var (
	errorCase = errors.New("some error")
	existData = user.User{
		FullName: "John Doe",
		Username: "johndoe",
		Password: "password",
		RoleID:   1,
	}
	nilData = user.User{}
)

func TestUseCaseShowAll(t *testing.T) {
	type fields struct {
		R user.RepositoryInterface
	}

	successCase := []user.User{}
	for i := 0; i < 5; i++ {
		successCase = append(successCase, existData)
	}

	mockRepository := mocks.NewRepositoryInterface(t)
	mockRepository.EXPECT().ShowAll().Return(successCase, nil).Once()
	mockRepository.EXPECT().ShowAll().Return(nil, errorCase).Once()

	tests := []struct {
		name    string
		fields  fields
		want    []user.User
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
			u := &user.UseCase{
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
		R user.RepositoryInterface
	}
	type args struct {
		user *user.User
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
				user: &existData,
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				user: &nilData,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user.UseCase{
				R: tt.fields.R,
			}
			if err := u.Create(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCaseShow(t *testing.T) {
	type fields struct {
		R user.RepositoryInterface
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
		want    *user.User
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
			u := &user.UseCase{
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
