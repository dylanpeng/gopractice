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
