package diary

type Diary struct {
	Id         id
	TalentName talentName
	GroupName  groupName
}

func NewDiary(id id, talentName talentName, groupName groupName) *Diary {
	return &Diary{
		Id:         id,
		TalentName: talentName,
		GroupName:  groupName,
	}
}
