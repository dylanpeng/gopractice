package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	ID    int    `gorm:"column:user_id" json:"id"`
	Name  string `gorm:"column:full_name" json:"name"`
	Email string `json:"email"`
}

func main() {
	// 测试指针切片 []*User
	queryData1 := []map[string]interface{}{
		{"user_id": 1, "full_name": "Alice", "email": "alice@example.com"},
		{"id": 2, "name": "Bob", "email": "bob@example.com"}, // 测试 json 标签
	}
	var ptrUsers []*User
	if err := AssignQueryResult(queryData1, &ptrUsers); err != nil {
		panic(err)
	}
	fmt.Printf("Pointer Slice: %+v\n", ptrUsers[0]) // &{ID:1 Name:Alice Email:alice@example.com}

	// 测试非指针切片 []User
	queryData2 := []map[string]interface{}{
		{"user_id": 3, "full_name": "Charlie", "email": "charlie@example.com"},
	}
	var structUsers []User
	if err := AssignQueryResult(queryData2, &structUsers); err != nil {
		panic(err)
	}
	fmt.Printf("Struct Slice: %+v\n", structUsers[0]) // {ID:3 Name:Charlie Email:charlie@example.com}
}

func AssignQueryResult(data []map[string]interface{}, modelsSlice interface{}) error {
	slicePtrValue := reflect.ValueOf(modelsSlice)
	if slicePtrValue.Kind() != reflect.Ptr || slicePtrValue.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("modelsSlice must be a pointer to a slice")
	}

	sliceValue := slicePtrValue.Elem()
	elemType := sliceValue.Type().Elem() // 获取切片元素类型（可能是结构体或指针）

	// 判断切片元素是否为指针类型
	isPtr := elemType.Kind() == reflect.Ptr
	var structType reflect.Type
	if isPtr {
		structType = elemType.Elem() // 获取指针指向的结构体类型
		if structType.Kind() != reflect.Struct {
			return fmt.Errorf("slice element must be a struct or pointer to struct")
		}
	} else {
		structType = elemType // 直接使用结构体类型
		if structType.Kind() != reflect.Struct {
			return fmt.Errorf("slice element must be a struct")
		}
	}

	for _, row := range data {
		// 创建新的结构体实例（如果是指针类型需要额外处理）
		newStructValue := reflect.New(structType).Elem() // 结构体实例

		// 遍历结构体所有字段
		for i := 0; i < newStructValue.NumField(); i++ {
			field := structType.Field(i)
			fieldValue := newStructValue.Field(i)

			// 获取列名逻辑（与之前相同）
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
			if columnName == "" {
				jsonTag := field.Tag.Get("json")
				if jsonTag != "" {
					columnName = strings.Split(jsonTag, ",")[0]
				} else {
					columnName = field.Name
				}
			}

			// 赋值逻辑（与之前相同）
			if val, ok := row[columnName]; ok && val != nil {
				rv := reflect.ValueOf(val)
				if rv.Type().ConvertibleTo(fieldValue.Type()) {
					fieldValue.Set(rv.Convert(fieldValue.Type()))
				} else {
					continue
				}
			}
		}

		// 根据切片类型决定添加指针还是结构体
		if isPtr {
			sliceValue.Set(reflect.Append(sliceValue, newStructValue.Addr())) // 添加指针
		} else {
			sliceValue.Set(reflect.Append(sliceValue, newStructValue)) // 添加结构体
		}
	}

	return nil
}
