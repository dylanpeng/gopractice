package global

import (
	"fmt"
	"gopractice/projects/gormshardingdemo/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

var (
	DB *gorm.DB
)

func InitDB() {
	dsn := "manager:123qweasd@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("init db fail. | err: %s", err)
		return
	}

	db.Use(sharding.Register(sharding.Config{
		ShardingKey:         "id",
		NumberOfShards:      4,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, &entity.Sharding{}))

	DB = db
}
