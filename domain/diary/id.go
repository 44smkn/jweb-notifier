package diary

type id string

func newId(diaryId string) (id, error) {
	// 桁数などでvalidataionをかける
	return id(diaryId), nil
}
