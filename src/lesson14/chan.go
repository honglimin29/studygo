package main

import "fmt"

func recv(in chan int) {
	ret := <-in
	fmt.Printf("接收成功，值为：%d\n", ret)
}

func chanDemo1() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	//发送0-99的数据
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	//从ch1接收数据并平方后发送给ch2
	go func() {
		for {
			data, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- data * data
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func square(out chan<- int, in <-chan int) {
	for {
		data, ok := <-in
		if !ok {
			break
		}
		out <- data * data
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func selectDemo() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

func main() {

	//ch := make(chan int)
	//go recv(ch)
	//ch <- 10
	//fmt.Println("发送成功")

	//chanDemo1()

	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go square(ch2, ch1)
	printer(ch2)

	selectDemo()

}
