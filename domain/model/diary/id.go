package diary

type Id struct {
	id string
}

func NewId(id string) (*Id, error) {
	// 桁数などでvalidataionをかける
	return &Id{id}, nil
}
