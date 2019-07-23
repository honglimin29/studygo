package main

import (
	"fmt"
	"sort"
)

//切片的定义
func defSlice() {
	var a []string
	var b = []int{}
	var c = []bool{true, false}
	var d = []bool{true, false}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)
	//fmt.Println(c == d) //不能对切片进行 == 比较，只能与nil比较
}

//基于数组定义切片
func sliceDemo1() {
	a := [5]int{55, 56, 57, 58, 59}
	s1 := a[2:]
	fmt.Printf("切片s1：%v ,len=%d,cap=%d,ptr=%p,type=%T", s1, len(s1), cap(s1), s1, s1)
}

//基于数组切片再切片
func sliceDemo2() {
	cityArray := [6]string{"北京", "上海", "广州", "深圳", "重庆", "成都"}
	fmt.Printf("value=%v,type=%T,len=%d,cap=%d\n", cityArray, cityArray, len(cityArray), cap(cityArray))
	s1 := cityArray[1:]
	fmt.Printf("value=%v,type=%T,len=%d,cap=%d\n", s1, s1, len(s1), cap(s1))
	s2 := s1[1:]
	fmt.Printf("value=%v,type=%T,len=%d,cap=%d\n", s2, s2, len(s2), cap(s2))
}

//切片的赋值拷贝
func sliceDemo3() {
	s1 := make([]int, 3, 3)
	s2 := s1
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
}

//切片的追加
func sliceDemo4() {
	var numSlice = []int{}
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("value=%v,len=%d,cap=%d,ptr=%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}

func sliceDemo5() {
	var citySlice = []string{}
	citySlice = append(citySlice, "北京")       //单个追加
	citySlice = append(citySlice, "上海", "广州") //多个追加
	c1 := []string{"深圳", "杭州"}
	citySlice = append(citySlice, c1...)
	fmt.Println(citySlice)
}

//切片的复制
func sliceDemo6() {
	s1 := make([]int, 3)
	s2 := make([]int, 3)
	copy(s2, s1)
	s1[0] = 100
	fmt.Printf("value=%v,len=%d,cap=%d,ptr=%p\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("value=%v,len=%d,cap=%d,ptr=%p\n", s2, len(s2), cap(s2), s2)

}

//切片删除元素
func sliceDemo7() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s1)
	s1 = append(s1[:2], s1[3:]...)
	fmt.Println(s1)
}

func Test1() {
	s1 := make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		fmt.Sprintf("%v", i)
		s1 = append(s1, fmt.Sprintf("%v", i))
	}
	fmt.Println(s1) //[     0 1 2 3 4 5 6 7 8 9]
	fmt.Printf("%d", len(s1))
}

//请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序
func Test2() {
	var a = [...]int{3, 7, 8, 9, 1}
	b := a[:]
	sort.Ints(b)
	fmt.Println(b)
}

func main() {
	//defSlice()
	//sliceDemo1()
	//sliceDemo2()
	//sliceDemo3()
	//sliceDemo4()
	//sliceDemo5()
	//sliceDemo6()
	//sliceDemo7()
	//Test1()
	Test2()
}
