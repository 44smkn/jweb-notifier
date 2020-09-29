package entity

import "jweb-notifier/domain/model/vo"

type Diary struct {
	Id         vo.DiaryId
	TalentName vo.TalentName
	GroupName  vo.GroupName
}
