package vo

type DiaryId struct {
	id string
}

func NewDiaryId(id string) (*DiaryId, error) {
	// 桁数などでvalidataionをかける
	return &DiaryId{id}, nil
}
