package diary

type Diary struct {
	id         Id
	talentName TalentName
	groupName  GroupName
}

func NewDiary(id Id, talentName TalentName, groupName GroupName) *Diary {
	return &Diary{
		id:         id,
		talentName: talentName,
		groupName:  groupName,
	}
}
