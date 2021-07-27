package main

import (
	"fmt"
	"gopractice/projects/gormdemo/table"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	dsn := "manager:123qweasd@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("init db fail. err: %s\n", err)
		return
	}

	AddDemo(db)

	demo1 := &table.Demo{}

	err = db.Model(&table.Demo{}).First(demo1).Error

	if err != nil {
		fmt.Printf("get demo fail. err: %s\n", err)
		return
	}

	demo2 := &table.Demo{}

	err = db.Table("demo").Where("id=?", 1).First(demo2).Error

	if err != nil {
		fmt.Printf("get demo fail. err: %s\n", err)
		return
	}
}

func AddDemo(db *gorm.DB) {
	now := time.Now()

	demo := &table.Demo{
		Name:      "test",
		Type:      1,
		Desc:      "描述",
		TimeStamp: now,
		TimeDate:  now,
		TimeDate2: &now,
	}

	err := db.Create(demo).Error

	if err != nil {
		fmt.Printf("create demo fail. err: %s\n", err)
		return
	}
}
