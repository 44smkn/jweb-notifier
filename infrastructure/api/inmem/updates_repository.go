package inmem

import (
	dd "jweb-notifier/domain/diary"
	du "jweb-notifier/domain/update"
	"log"
	"time"
)

type UpdatesRepository struct{}

func (u *UpdatesRepository) Fetch() *du.Updates {
	dt := du.NewDateTime(time.Date(2020, time.October, 10, 18, 0, 0, 0, time.Local))
	id, err := dd.NewId("59")
	if err != nil {
		log.Fatalln("diary idの生成に失敗しました")
	}
	tn, err := dd.NewTalentName("川島如恵留")
	if err != nil {
		log.Fatalln("talent nameの生成に失敗しました")
	}
	gn := dd.NewGroupName("TravisJapan")
	diaries := []dd.Diary{
		*dd.NewDiary(id, tn, gn),
	}
	updates := du.NewUpdates(dt, diaries)
	return updates
}
