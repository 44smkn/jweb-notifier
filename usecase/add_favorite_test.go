package usecase_test

import (
	"jweb-notifier/presentation/param"
	"jweb-notifier/usecase"
	"strings"
	"testing"
)

func TestAddFavorite(t *testing.T) {
	t.Parallel()
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
			if err != nil && !strings.Contains(err.Error(), tt.want.Error()) {
				t.Errorf("want: %v  got: %v", tt.want.Error(), err.Error())
			}
		})
	}
}
