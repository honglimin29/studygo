package main

import "fmt"

type Person struct {
	name string
	city string
	age  int
}

func NewPerson(name string, city string, age int) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}

func (p Person) dream() {
	fmt.Printf("%s的梦想是学号GO语言!!", p.name)
}

func main() {

	var p1 Person
	fmt.Printf("type:%T,value:%#v\n", p1, p1)

	p2 := new(Person)
	p2.name = "沙河小王子"
	p2.city = "广州"
	p2.age = 22
	fmt.Printf("type:%T,value:%#v\n", p2, p2)

	p3 := Person{
		name: "honglimin",
		city: "上海",
		age:  22,
	}
	fmt.Printf("type:%T,value:%#v\n", p3, p3)

	p4 := &Person{
		name: "tangqige",
		city: "北京",
		age:  23,
	}
	fmt.Printf("type:%T,value:%#v\n", p4, p4)

	p5 := NewPerson("沙河小王子", "海南", 22)
	fmt.Printf("type:%T,value:%#v\n", p5, p5)
	p5.dream()

}
