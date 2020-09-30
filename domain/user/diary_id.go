package user

type diaryId string

func NewDiaryId(id string) (diaryId, error) {
	// 桁数などでvalidataionをかける
	return diaryId(id), nil
}
