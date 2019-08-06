package main

import (
	"fmt"
	"strings"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

var (
	coins        = 50
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	left := coins
	distributionCoin := func(name string) {
		coin := 0
		upperName := strings.ToUpper(name)
		if ret := strings.Count(upperName, "E"); ret > 0 {
			coin += 1 * ret
		}
		if ret := strings.Count(upperName, "I"); ret > 0 {
			coin += 2 * ret
		}
		if ret := strings.Count(upperName, "O"); ret > 0 {
			coin += 3 * ret
		}
		if ret := strings.Count(upperName, "U"); ret > 0 {
			coin += 4 * ret
		}
		distribution[name] = coin
		left -= coin
	}
	for _, name := range users {
		distributionCoin(name)
	}

	return left
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println(distribution)
}
