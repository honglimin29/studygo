/**
Go语言基础之变量和常量
*/

package main

import "fmt"

/*
变量的声明
*/

var m int = 100

const pi = 3.1415

const (
	_  = iota
	kb = 1 << (10 * iota) // 1 * 2的10次方
	mb = 1 << (10 * iota) // 1 * 2的20次方
	gb = 1 << (10 * iota) // 1 * 2的30次方
	tb = 1 << (10 * iota) // 1 * 2的40次方
	pb = 1 << (10 * iota) // 1 * 2的50次方
)

func foo() (int, string) {
	return 22, "honglimin"
}

func main() {

	var name string
	name = "洪李敏"
	n := 100

	fmt.Println(m, n)
	fmt.Println(pi)
	fmt.Println(name)

	fmt.Println("kb=", kb)
	fmt.Println("mb=", mb)
	fmt.Println("gb=", gb)
	fmt.Println("tb=", tb)
	fmt.Println("pb=", pb)

	//匿名变量
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

}
