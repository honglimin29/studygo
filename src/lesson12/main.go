package main

import "fmt"

type Sayer interface {
	say()
}

type Animal struct {
	name string
}

type Dog struct {
	*Animal
}

type Cat struct {
	*Animal
}

func (d *Dog) say() {
	fmt.Printf("%s,旺旺旺!!!\n", d.name)
}

func (c *Cat) say() {
	fmt.Printf("%s,喵喵喵!!!\n", c.name)
}

func show(a interface{}) {
	fmt.Printf("type:%T,value:%v\n", a, a)
	v, ok := a.(string)
	if ok {
		fmt.Printf("类型为string,%s", v)
	}
}

func main() {
	var s Sayer
	wangcai := &Dog{&Animal{"旺财"}}
	s = wangcai
	s.say()

	xiaoxin := &Cat{&Animal{"小新"}}
	s = xiaoxin
	s.say()

	show("hello world")
}
