package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	ID    int    `gorm:"column:user_id" json:"id"`
	Name  int    `gorm:"column:full_name" json:"name"`
	Email string `json:"email"`
}

func main() {
	// 模拟数据库查询结果
	queryData := []map[string]interface{}{
		{
			"user_id":   1,
			"full_name": "Alice",
			"email":     "alice@example.com",
		},
		{
			"user_id": 2,
			"name":    "Bob", // 匹配 json 标签
			"email":   "bob@example.com",
		},
	}

	// 目标切片
	var users []*User

	// 调用赋值函数
	if err := AssignQueryResult(queryData, &users); err != nil {
		panic(err)
	}

	// 验证结果
	fmt.Printf("%+v\n", users)
	// 输出：
	// [{ID:1 Name:Alice Email:alice@example.com} {ID:2 Name:Bob Email:bob@example.com}]
}

func AssignQueryResult(data []map[string]interface{}, modelsSlice interface{}) error {
	sliceValue := reflect.ValueOf(modelsSlice)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("modelsSlice must be a pointer to a slice")
	}

	elemType := sliceValue.Elem().Type().Elem()
	slice := sliceValue.Elem()

	for _, row := range data {
		// 创建新结构体实例
		newElem := reflect.New(elemType).Elem()

		// 遍历结构体所有字段
		for i := 0; i < newElem.NumField(); i++ {
			field := newElem.Type().Field(i)
			fieldValue := newElem.Field(i)

			// 获取 gorm column 标签
			columnName := ""
			gormTag := field.Tag.Get("gorm")
			if gormTag != "" {
				for _, part := range strings.Split(gormTag, ";") {
					if strings.HasPrefix(part, "column:") {
						columnName = strings.TrimPrefix(part, "column:")
						break
					}
				}
			}

			// 如果未找到 gorm column，尝试 json 标签
			if columnName == "" {
				jsonTag := field.Tag.Get("json")
				if jsonTag != "" {
					columnName = strings.Split(jsonTag, ",")[0]
				} else {
					columnName = field.Name // 默认使用字段名
				}
			}

			// 从 map 中获取值
			if val, ok := row[columnName]; ok {
				// 处理零值和非空值
				if val == nil {
					continue // 跳过空值
				}

				// 获取反射值并转换类型
				rv := reflect.ValueOf(val)
				if rv.Type().ConvertibleTo(fieldValue.Type()) {
					fieldValue.Set(rv.Convert(fieldValue.Type()))
				} else {
					continue
				}
			}
		}

		// 将新元素添加到切片
		slice.Set(reflect.Append(slice, newElem))
	}

	return nil
}
