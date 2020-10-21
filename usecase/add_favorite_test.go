package usecase_test

import (
	dd "jweb-notifier/domain/diary"
	du "jweb-notifier/domain/user"
	"jweb-notifier/infrastructure/persistence/inmem"
	"jweb-notifier/presentation"
	"jweb-notifier/usecase"
	"strings"
	"testing"
)

func TestAddFavorite(t *testing.T) {
	t.Parallel()

	repo := &inmem.UserRepository{}
	err := repo.Register(createUser("44smkn@gmail.com", "password"))
	if err != nil {
		t.Errorf("failed to register user")
	}

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
			f := &presentation.AddFavoriteParam{
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

func createUser(id, password string) *du.User {
	userId, _ := du.NewId(id)
	userPassword, _ := du.NewPassword(password)
	return du.NewUser(userId, userPassword, []dd.Id{})
}
