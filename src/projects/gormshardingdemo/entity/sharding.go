package entity

import "time"

type Sharding struct {
	ID          int64  `gorm:"primaryKey"`
	Name        string `gorm:"column:name"`
	CreatedTime time.Time
	UpdatedTime time.Time
}

func (e *Sharding) TableName() string {
	return "sharding"
}
