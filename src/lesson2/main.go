package main

import (
	"fmt"
	"math"
	"strings"
)

/*
遍历某字符串
*/
func traverseString() {
	s1 := "hello"
	for i := 0; i < len(s1); i++ { //bytes
		fmt.Printf("%v(%c)", s1[i], s1[i])
	}
	fmt.Println()
	s2 := "hello小王子"
	for _, v := range s2 { //rune
		fmt.Printf("%v(%c)", v, v)
	}
	fmt.Println()
}

/*
修改某字符串
*/
func changeString() {
	s1 := "big"
	byte1 := []byte(s1)
	byte1[0] = 'p'
	fmt.Println(s1)
	fmt.Println(string(byte1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(s2)
	fmt.Println(string(runeS2))

}

/*
编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型
*/
func Test1() {

	//定义变量
	var a int = 10
	var b float32 = math.Pi
	var c bool = true
	var d string = "沙河小王子"

	//打印
	fmt.Printf("a变量的类型 ：%T , 变量值 ：%d\n", a, a)
	fmt.Printf("b变量的类型 ：%T , 变量值 ：%.2f\n", b, b)
	fmt.Printf("c变量的类型 ：%T , 变量值 ：%t\n", c, c)
	fmt.Printf("d变量的类型 ：%T , 变量值 ：%s\n", d, d)
}

/*
编写代码统计出字符串"hello沙河小王子"中汉字的数量。
*/
func Test2(contents string) {
	letter := "abcdefghijklmnopqrstuvwxyz"
	letter = letter + strings.ToUpper(letter)
	nums := "0123456789"

	var (
		total        int = 0
		letter_count int = 0
		num_count    int = 0
		other_count  int = 0
	)

	for _, content := range contents {

		switch {
		case strings.ContainsRune(letter, content) == true:
			letter_count += 1
		case strings.ContainsRune(nums, content) == true:
			num_count += 1
		default:
			other_count += 1
		}

		total++
	}

	fmt.Printf("统计结果: 总字数=%d , 字母数=%d , 数字=%d , 汉字=%d", total, letter_count, num_count, other_count)

}

func main() {

	Test1()
	Test2("hello沙河sss小王子123,,")

	fmt.Println()

	traverseString()

	changeString()

}
