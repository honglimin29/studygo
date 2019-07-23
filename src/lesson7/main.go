package main

import "fmt"

func pointerDemo1() {
	a := 10
	b := &a
	a = 100
	fmt.Printf("变量a的值:%v,变量a的内存地址:%v,变量a的类型:%T\n", a, &a, a)
	fmt.Printf("变量b的值:%v,变量b的内存地址:%v,变量b的类型:%T,变量b所指向的值:%v\n", b, &b, b, *b)
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func pointerDemo2() {
	a := 10
	modify1(a)
	fmt.Println(a)
	b := 10
	modify2(&b)
	fmt.Println(b)
}

//new的用法,返回对应类型的指针类型
func pointerDemo3() {
	a := new(int)
	b := new(bool)
	c := new(string)
	d := new(byte)
	e := new(*[3]int)
	fmt.Printf("a=%T,b=%T,c=%T,d=%T,e=%T", a, b, c, d, e)
}

//make的用法，分配内存，返回对应类型的本身
func pointerDemo4() {
	a := make([]int, 3, 5)
	fmt.Printf("value=%v,type=%T\n", a, a)
	b := make(map[string]string)
	b["name"] = "洪李敏"
	fmt.Printf("value=%v,type=%T\n", b, b)
}

func main() {
	//pointerDemo1()
	//pointerDemo2()
	//pointerDemo3()
	pointerDemo4()
}
