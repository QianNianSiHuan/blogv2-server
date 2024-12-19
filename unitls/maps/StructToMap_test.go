package maps_test

import (
	"blogv2/unitls/maps"
	"fmt"
	"testing"
)

// TestStructToMap 测试 StructToMap 函数.
func TestStructToMap(t *testing.T) {
	type successTestStruct struct {
		String string `tag-string:"string"`
		Int    int    `tag-number:"int"`
		Bool   bool   `tag-bool:"bool"`
	}
	var test = successTestStruct{
		String: "测试",
		Int:    10086,
		Bool:   false,
	}
	var testMap map[string]any
	var err error
	if testMap, err = maps.StructToMap(test, "tag-string"); testMap["string"] != "测试" {
		fmt.Println(err)
		t.Fatal("字符串结构体转map失败")
	} else {
		fmt.Println(testMap)
		t.Log("字符串结构体转map成功")
	}
	if testMap, err = maps.StructToMap(test, "tag-number"); testMap["int"] != 10086 {
		fmt.Println(err)
		t.Fatal("整型结构体转map失败")
	} else {
		fmt.Println(testMap)
		t.Log("整型结构体转map成功")
	}
	if testMap, err = maps.StructToMap(test, "tag-bool"); testMap["bool"] != false {
		fmt.Println(err)
		t.Fatal("布尔结构体转map失败")
	} else {
		fmt.Println(testMap)
		t.Log("布尔结构体转map成功")
	}
}
