package vo

type DiaryId string

func NewDiaryId(id string) (DiaryId, error) {
	// 桁数などでvalidataionをかける
	return DiaryId(id), nil
}
