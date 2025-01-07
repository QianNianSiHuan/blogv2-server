package maps

import (
	"errors"
	"reflect"
)

func StructToMap(data any, tagName string) (mp map[string]any, err error) {
	mp = make(map[string]any)
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		err = errors.New("传入转化的对象非结构体")
		return
	}
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i)
		tag := v.Type().Field(i).Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}
		if val.Kind() == reflect.Ptr {
			if val.IsNil() {
				continue
			}
			mp[tag] = val.Elem().Interface()
			continue
		}
		mp[tag] = val.Interface()
	}
	return
}
