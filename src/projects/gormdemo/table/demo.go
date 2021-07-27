package table

import (
	"time"
)

type Demo struct {
	ID        int64      `gorm:"column:id;primaryKey" json:"id"`
	Name      string     `json:"name"`
	Type      int        `json:"type"`
	Desc      string     `json:"desc"`
	TimeStamp time.Time  `json:"time_stamp"`
	TimeDate  time.Time  `json:"time_date"`
	TimeDate2 *time.Time `json:"time_date2"`
}

func (e *Demo) TableName() string {
	return "demo"
}
