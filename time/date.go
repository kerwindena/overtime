package time

import (
	"time"
)

type Date struct {
	time time.Time
	day  int
}

func newDate(day int) *Date {
	return &Date{
		time: time.Unix(0, 0).AddDate(0, 0, day),
		day:  day,
	}
}

func (date *Date) String() string {
	return date.time.String()
}

func (date *Date) Day() int {
	return date.day
}
