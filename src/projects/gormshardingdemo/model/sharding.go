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

func (m *shardingModel) GetDatas(ids []int64) (datas []*entity.Sharding, err error) {
	datas = make([]*entity.Sharding, 0)
	//datas := make([]*entity.Sharding, 0)

	db := global.DB
	for _, id := range ids {
		data := new(entity.Sharding)
		e := db.Model(&entity.Sharding{}).Where("id", id).First(data).Error

		if e == nil {
			datas = append(datas, data)
		} else if err == nil {
			err = e
		}
	}

	//if err == nil {
	//	data = datas[0]
	//}

	return
}

func (m *shardingModel) Update(data *entity.Sharding) (err error) {
	db := global.DB
	err = db.Exec("UPDATE sharding SET name = ? WHERE id = ?", "name", int64(1687323722520)).Error

	return
}
