package update

import (
	"time"
)

type DateTime struct {
	val time.Time
}

func NewDateTime(dateTime time.Time) DateTime {
	return DateTime{val: dateTime}
}
