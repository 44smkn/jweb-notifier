package inmem

import "jweb-notifier/domain/model/entity"

type Updater interface {
	Fetch() *entity.Updates
}
