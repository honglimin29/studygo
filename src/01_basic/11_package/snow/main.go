package main

import (
	"fmt"
	clac2 "github.com/honglimin29/studygo/src/basic/11_package/clac"
)

func main() {
	var x, y int = 20, 10

	ret1 := clac2.Adder(x, y)
	fmt.Println(ret1)

	ret2 := clac2.Suber(x, y)
	fmt.Println(ret2)

	ret3 := clac2.Taker(x, y)
	fmt.Println(ret3)

	ret4 := clac2.Divide(x, y)
	fmt.Println(ret4)
}
