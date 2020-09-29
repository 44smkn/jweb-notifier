package diary

type groupName string

func NewGroupName(name string) groupName {
	return groupName(name)
}
