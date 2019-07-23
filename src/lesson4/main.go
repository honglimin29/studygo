package main

import (
	"fmt"
	"time"
)

func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("成绩A")
	} else if score >= 75 {
		fmt.Println("成绩B")
	} else {
		fmt.Println("成绩C")
	}
}

func ifDemo2() {
	if score := 80; score >= 90 {
		fmt.Println("成绩A")
	} else if score >= 75 {
		fmt.Println("成绩B")
	} else {
		fmt.Println("成绩C")
	}
}

//死循环
func forDemo1() {
	for {
		fmt.Println("helloworld!!")
		time.Sleep(1 * time.Second)
	}
}

//for循环常规写法
func forDemo2() {
	n := 1
	for i := 0; i < 10; i++ {
		fmt.Println(n)
		n++
	}
}

// for循环省略第一个声明
func forDemo3() {
	i := 1
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

// for循环实现 while
func forDemo4() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

// case
func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入")
	}
}

// case demo2
func switchDemo2() {
	num := 3
	switch num {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 0, 2, 4, 6, 8:
		fmt.Println("偶数")
	}
}

func switchDemo3() {
	score := 64
	switch {
	case score > 80:
		fmt.Println("你很优秀!!")
	case score > 60 && score <= 80:
		fmt.Println("还需努力!!")
		fallthrough // 兼容C用法，匹配当前和下一个case
	default:
		fmt.Println("放弃吧!!")
	}
}

// goto DEMO
func gotoDemo() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakFlag
			}
		}
	}

breakFlag:
	fmt.Println("结束循环")
}

//编写代码打印9*9乘法表。

func calc() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {

	//ifDemo1()
	//ifDemo2()
	//forDemo1()
	//forDemo2()
	//forDemo3()
	//forDemo4()
	//switchDemo1()
	//switchDemo2()
	//switchDemo3()
	//gotoDemo()
	calc()

}
