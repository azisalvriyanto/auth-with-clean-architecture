package auth

import (
	"auth-with-clean-architecture/internal/auth"
	"auth-with-clean-architecture/internal/auth/mocks"
	"errors"
	"reflect"
	"testing"
)

func TestUseCaseLogin(t *testing.T) {
	type fields struct {
		R auth.RepositoryInterface
	}
	type args struct {
		payload *auth.Payload
	}

	errorCase := errors.New("some error")
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
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywiaWF0IjoxNjgzMjY4NjMyfQ.s6YHhgiPlYcZHCXqaIlHKy784SSGUSuvVHxutNo6HVY",
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
