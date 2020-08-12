package vo

type DiaryId struct {
	id string
}

func NewDiaryId(id string) (*Id, error) {
	// 桁数などでvalidataionをかける
	return &Id{id}, nil
}
