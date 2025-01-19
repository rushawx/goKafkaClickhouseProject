package stats

import "time"

// Stats struct
type Stats struct {
	Day   time.Time `gorm:"column:day"`
	Total int       `gorm:"column:total"`
}
