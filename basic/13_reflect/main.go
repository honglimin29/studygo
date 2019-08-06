package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name  string `json:"name"`
	score int    `json:"score"`
}

func main() {

	var s1 = student{
		name:  "沙河小王子",
		score: 22,
	}

	t := reflect.TypeOf(s1)
	v := reflect.ValueOf(s1)
	fmt.Println(t.Name(), t.Kind())
	fmt.Println(v.Kind())
	//遍历结构体所有字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("index:%v,name:%v,type:%v,tag:%v\n", field.Index, field.Name, field.Type, field.Tag.Get("json"))
	}
	//通过指定字段查找结构体字段
	if fieldItem, ok := t.FieldByName("score"); ok {
		fmt.Printf("index:%v,name:%v,type:%v,tag:%v\n", fieldItem.Index, fieldItem.Name, fieldItem.Type, fieldItem.Tag.Get("json"))
	}

}
