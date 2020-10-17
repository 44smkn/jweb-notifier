package usecase_test

import (
	"jweb-notifier/domain/diary"
	"jweb-notifier/domain/user"
	"jweb-notifier/presentation/param"
	"jweb-notifier/usecase"
	"strings"
	"testing"
)

func TestAddFavorite(t *testing.T) {
	t.Parallel()

	inmem.Users = append(inmem.Users, createUser("44smkn@gmail.com", "password"))

	tests := []struct {
		name    string
		userId  string
		diaryId string
		want    error
	}{
		{
			name:    "normal case",
			userId:  "44smkn@gmail.com",
			diaryId: "fatatgaga00",
			want:    nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := &param.AddFavoriteInput{
				UserId:  tt.userId,
				DiaryId: tt.diaryId,
			}
			err := usecase.AddFavorite(f)
			if err != tt.want && !strings.Contains(err.Error(), tt.want.Error()) {
				t.Errorf("want: %v  got: %v", tt.want.Error(), err.Error())
			}
		})
	}
}

func createUser(id, password string) user.User {
	userId, _ := user.NewId(id)
	userPassword, _ := user.NewPassword(password)
	return user.User{
		id: userId,
		password: userPassword,
		favorite: []diary.Id{}
	}
}
