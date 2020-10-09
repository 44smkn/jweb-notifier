package param

type Favorite struct {
	UserId  string
	DiaryId string
}

func NewFavorite(userId, diaryId string) Favorite {
	return Favorite{
		UserId:  userId,
		DiaryId: diaryId,
	}
}
