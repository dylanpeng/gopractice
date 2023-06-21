package logic

import (
	"fmt"
	"gopractice/projects/gormshardingdemo/entity"
	"gopractice/projects/gormshardingdemo/model"
	"time"
)

var Logic = &logic{}

type logic struct {
}

func (l *logic) AddSharding() {

	for i := 0; i < 100; i++ {
		id := time.Now().UnixMilli()
		data := &entity.Sharding{
			ID:          id,
			Name:        "test",
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}

		err := model.Sharding.Add(data)

		if err != nil {
			fmt.Printf("sharding add fail. | err: %s\n", err)
			continue
		}

		fmt.Printf("add sharding success. data: %+v\n", *data)

		data1, _ := model.Sharding.GetData(id)

		fmt.Printf("get sharding success. data: %+v\n", *data1)
	}

}

func (l *logic) GetSharding() {
	datas, err := model.Sharding.GetDatas([]int64{1687323722499, 98, 1687323722513, 1687323722520})

	if err != nil {
		fmt.Printf("sharding get fail. | err: %s\n", err)
		return
	}

	fmt.Printf("sharding get success. | datas: %s\n", datas)
}

func (l *logic) Update() {
	err := model.Sharding.Update(nil)

	if err != nil {
		fmt.Printf("sharding get fail. | err: %s\n", err)
		return
	}
}
