package entity

import (
	"fmt"
	"time"
)

type Sharding struct {
	ID          int64  `gorm:"primaryKey"`
	Name        string `gorm:"column:name"`
	CreatedTime time.Time
	UpdatedTime time.Time
}

func (e *Sharding) String() string {
	return fmt.Sprintf("%+v", *e)
}

func (e *Sharding) TableName() string {
	return "sharding"
}
