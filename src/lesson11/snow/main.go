package main

import (
	"fmt"
	"github.com/honglimin29/studygo/src/lesson11/clac"
)

func main() {
	var x, y int = 20, 10

	ret1 := clac.Adder(x, y)
	fmt.Println(ret1)

	ret2 := clac.Suber(x, y)
	fmt.Println(ret2)

	ret3 := clac.Taker(x, y)
	fmt.Println(ret3)

	ret4 := clac.Divide(x, y)
	fmt.Println(ret4)
}
