package stats

import "time"

type stats struct {
	Day   time.Time `gorm:"column:day"`
	Total int       `gorm:"column:total"`
}
