package utils

import (
	"reflect"
)

// @function:StructToMap
// @description:利用反射将结构体转化为map
// @param:interface{}
// @return:map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

// @function:Pointer
// @description: 返回变量对应指针
// @param: T any
// @return: *T
func Pointer[T any](in T) (out *T) {
	return &in
}
