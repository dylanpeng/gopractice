package common

import "encoding/json"

// 对象转换
func ConvertStruct(a interface{}, b interface{}) error {
	err := convertStruct(a, b)
	return err
}

func convertStruct(a interface{}, b interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, b)
	if err != nil {
		return err
	}
	return nil
}
