package auth

import (
	"auth-with-clean-architecture/internal/auth"
	"auth-with-clean-architecture/internal/auth/mocks"
	"auth-with-clean-architecture/internal/user"
	"errors"
	"reflect"
	"testing"
)

var (
	tokenSigned = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywiaWF0IjoxNjgzMjY4NjMyfQ.s6YHhgiPlYcZHCXqaIlHKy784SSGUSuvVHxutNo6HVY"
	errorCase   = errors.New("some error")
)

func TestUseCaseLogin(t *testing.T) {
	type fields struct {
		R auth.RepositoryInterface
	}
	type args struct {
		payload *auth.Payload
	}

	existData := auth.Payload{
		Username: "username",
		Password: "password",
	}
	nilData := auth.Payload{}
	existDataResponse := auth.ProfileItemWithToken{
		ProfileItem: auth.ProfileItem{
			ID:       1,
			FullName: "John Doe",
			Username: "username",
		},
		Token: tokenSigned,
	}
	nilDataResponse := auth.ProfileItemWithToken{}

	mockRepository := mocks.NewRepositoryInterface(t)
	mockRepository.EXPECT().Login(&existData).Return(&existDataResponse, nil).Once()
	mockRepository.EXPECT().Login(&nilData).Return(&nilDataResponse, errorCase).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *auth.ProfileItemWithToken
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				payload: &existData,
			},
			want:    &existDataResponse,
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				payload: &nilData,
			},
			want:    &nilDataResponse,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &auth.UseCase{
				R: tt.fields.R,
			}
			got, err := u.Login(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseShowProfile(t *testing.T) {
	type fields struct {
		R auth.RepositoryInterface
	}
	type args struct {
		tokenSigned string
	}

	existDataResponse := user.User{
		FullName: "John Doe",
		Username: "johndoe",
		Password: "password",
		RoleID:   1,
	}
	nilDataResponse := user.User{}

	mockRepository := mocks.NewRepositoryInterface(t)
	mockRepository.EXPECT().ShowProfile(tokenSigned).Return(&existDataResponse, nil).Once()
	mockRepository.EXPECT().ShowProfile("").Return(&nilDataResponse, errorCase).Once()

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
				tokenSigned: tokenSigned,
			},
			want:    &existDataResponse,
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				tokenSigned: "",
			},
			want:    &nilDataResponse,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &auth.UseCase{
				R: tt.fields.R,
			}
			got, err := u.ShowProfile(tt.args.tokenSigned)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.ShowProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.ShowProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
