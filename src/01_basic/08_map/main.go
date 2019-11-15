package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//map的基本使用
func mapDemo1() {
	//定义
	scoreMap := make(map[string]int, 8)
	scoreMap1 := map[string]int{
		"张三": 50,
		"李四": 99,
	}
	//新增
	scoreMap["张三"] = 50
	scoreMap["李四"] = 99
	scoreMap["王五"] = 70
	//修改
	if _, ok := scoreMap["张三"]; ok {
		scoreMap["张三"] = 88
	}
	//删除
	delete(scoreMap, "王五")
	fmt.Println(scoreMap)
	fmt.Println(scoreMap1)
}

//按照指定顺序遍历map
func mapDemo2() {
	rand.Seed(time.Now().UnixNano()) //初始化随机种子

	scoreMap := make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		val := rand.Intn(100)
		scoreMap[key] = val
	}

	keys := make([]string, 0)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s = %d\n", key, scoreMap[key])
	}
}

//map类型的切片
func mapDemo3() {
	mapSlice := make([]map[string]int, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d,value:%v\n", index, value)
	}
	mapScore := map[string]int{
		"张三": 88,
		"李四": 90,
	}
	mapSlice[0] = mapScore
	for index, value := range mapSlice {
		fmt.Printf("index:%d,value:%v\n", index, value)
	}
}

//切片类型的map
func mapDemo4() {
	sliceMap := make(map[string][]string)
	for key, value := range sliceMap {
		fmt.Printf("key:%s,value:%v\n", key, value)
	}
	if _, ok := sliceMap["广东省"]; !ok {
		sliceMap["广东省"] = []string{"广州市", "深圳市", "珠海市", "中山市", "东莞市"}
	}
	for key, value := range sliceMap {
		fmt.Printf("key:%s,value:%v\n", key, value)
	}

}

func main() {
	//mapDemo1()
	//mapDemo2()
	//mapDemo3()
	mapDemo4()
}
