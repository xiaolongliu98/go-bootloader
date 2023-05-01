package util

import (
	"fmt"
	"reflect"
)

// GetTypeName 返回值的类型名称，如果获取失败则返回错误
func GetTypeName(v interface{}) string {
	if v == nil {
		panic(fmt.Errorf("%v 类型错误", v))
	}
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		// 如果是指针类型，则获取指针指向的类型
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		// 如果不是结构体类型，则返回错误
		panic(fmt.Errorf("%v 类型错误", v))
	}
	return t.Name()
}
