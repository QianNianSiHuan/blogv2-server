package maps

import (
	"encoding/json"
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
		if tag == "" || tag == "-" || val.IsNil() {
			continue
		}
		if val.Kind() == reflect.Ptr {
			v1 := val.Elem().Interface()
			if val.Elem().Kind() == reflect.Slice {
				byteData, _ := json.Marshal(v1)
				mp[tag] = string(byteData)
			} else {
				mp[tag] = v1
			}
			continue
		}
		mp[tag] = val.Interface()
	}
	return
}
