package diary

type GroupName struct {
	val string
}

func NewGroupName(name string) GroupName {
	return GroupName{val: name}
}
