package vo

type DiaryId struct {
	id string
}

func NewDiaryId(id string) (*Id, error) {
	// 桁数などでvalidataionをかける
	return &Id{id}, nil
}

func (d *DiaryId) Equals(vo shared.ValueObject) bool {
	if d == vo {
		return true
	}
	if val, ok := vo.(DiaryId); !ok {
		return false
	}
	return d.id == val
}

func (d *DiaryId) GetValue() {
	return d.id
}