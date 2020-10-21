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

func (d DateTime) Format(layout string) string {
	return d.val.Format(layout)
}
