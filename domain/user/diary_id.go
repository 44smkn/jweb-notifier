package user

type diaryId string

func newDiaryId(id string) (diaryId, error) {
	// 桁数などでvalidataionをかける
	return diaryId(id), nil
}
