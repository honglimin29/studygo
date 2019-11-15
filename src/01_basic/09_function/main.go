package main

import (
	"errors"
	"fmt"
	"strings"
)

//函数的基本用法
func intSum(x, y int) (sum int) {
	sum = x + y
	return
}

//高阶函数
//函数作为参数
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x int, y int, op func(int, int) int) int {
	return op(x, y)
}

//函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		errs := errors.New("无法识别的操作符")
		return nil, errs
	}
}

//闭包函数
func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//defer
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

//defer面试题
func deferDemo(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//panic/recover
func funcA() {
	fmt.Println("funcA")
}

func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover in funcB")
		}
	}()
	panic("panic in funcB")
}

func funcC() {
	fmt.Println("funcC")
}

func main() {
	sum := intSum(10, 20)
	fmt.Println(sum)

	ret2 := calc(10, 20, add)
	fmt.Println(ret2)

	//匿名函数
	add := func(x, y int) int {
		return x + y
	}
	ret3 := add(10, 40)
	fmt.Println(ret3)

	//匿名函数自己执行
	add2 := func(x, y int) int {
		return x + y
	}(33, 33)
	fmt.Println(add2)

	//闭包调用
	ff := adder(10)
	fmt.Printf("type:%T\n", ff)
	fmt.Println(ff(10))
	fmt.Println(ff(20))

	var jpgFunc = makeSuffix(".jpg")
	txtFunc := makeSuffix(".txt")
	fmt.Println(jpgFunc("name"))
	fmt.Println(txtFunc("name"))

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	//经典面试题：结论：按顺序执行，遇defer关键字，先赋值，后延迟执行
	//a := 1
	//b := 2
	//defer deferDemo("1",a,deferDemo("10",a,b))
	//a = 0
	//defer deferDemo("2",a,deferDemo("20",a,b))
	//b = 1

	funcA()
	funcB()
	funcC()

}
