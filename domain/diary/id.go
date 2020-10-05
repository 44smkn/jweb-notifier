package diary

type Id struct {
	val string
}

func NewId(diaryId string) (Id, error) {
	// 桁数などでvalidataionをかける
	return Id{val: diaryId}, nil
}
