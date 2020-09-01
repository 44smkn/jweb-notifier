package vo

type DiaryId struct {
	id string
}

func NewDiaryId(id string) (*DiaryId, error) {
	// 桁数などでvalidataionをかける
	return &DiaryId{id}, nil
}

func (u *DiaryId) Equals(another *DiaryId) bool {
	return u.id == another.id
}
