package repository

import (
	"jweb-notifier/domain/model/entity"
	"jweb-notifier/domain/model/vo"
	"log"
)

type UpdatesRepository struct{}

func (u *UpdatesRepository) Fetch() *entity.Updates {
	diaryId, err := vo.NewDiaryId("1")
	if err != nil {
		log.Fatalln("failed to create diary id.")
	}
	return &entity.Updates{
		Date: "2020/09/16",
		Time: "15:00",
		Diaries: []entity.Diary{
			{
				Id:         diaryId,
				TalentName: vo.NoeruKawashima,
				GroupName:  vo.TravisJapan,
			},
		},
	}
}
