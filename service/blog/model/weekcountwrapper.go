package model

type WeekCount struct {
	Day   int64 `db:"day"`
	Count int64 `count:"count"`
}
