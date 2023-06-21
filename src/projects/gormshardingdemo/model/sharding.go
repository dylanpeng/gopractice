package model

import (
	"gopractice/projects/gormshardingdemo/entity"
	"gopractice/projects/gormshardingdemo/global"
	_ "gorm.io/sharding"
)

var Sharding = &shardingModel{}

type shardingModel struct {
}

func (m *shardingModel) Add(data *entity.Sharding) (err error) {
	db := global.DB
	err = db.Create(data).Error

	return
}

func (m *shardingModel) GetData(id int64) (data *entity.Sharding, err error) {
	data = &entity.Sharding{}
	//datas := make([]*entity.Sharding, 0)

	db := global.DB
	err = db.Model(&entity.Sharding{}).First(data, "id", id).Error

	//if err == nil {
	//	data = datas[0]
	//}

	return
}
