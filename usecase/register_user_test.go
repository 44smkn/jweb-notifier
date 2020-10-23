package usecase_test

import (
	"errors"
	"jweb-notifier/usecase"
	"strings"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		id       string
		password string
		want     error
	}{
		{
			name:     "normal case",
			id:       "44smkn@gmail.com",
			password: "abc234Gfu6s",
			want:     nil,
		},
		{
			name:     "invalid id",
			id:       "44smkn",
			password: "abc234Gfu6s",
			want:     errors.New("IDがメールアドレスの形式になっていません"),
		},
		{
			name:     "invalid password",
			id:       "44smkn@gmail.com",
			password: "1234567",
			want:     errors.New("英数字8文字以上ではありません"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			input := &usecase.RegisterUserParam{
				Id:       tt.id,
				Password: tt.password,
			}
			err := usecase.RegisterUser(input)
			if err != nil && !strings.Contains(err.Error(), tt.want.Error()) {
				t.Errorf("want: %v  got: %v", tt.want.Error(), err.Error())
			}
		})
	}

}
