package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id    int            `json:"id"`
	Name  string         `json:"name"`
	Age   int            `json:"age"`
	Score map[string]int `json:"score"`
}

type Class struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Students []*Student `json:"students"`
}

func NewClass(id int, name string) *Class {
	return &Class{
		Id:       id,
		Name:     name,
		Students: make([]*Student, 0),
	}
}

func (c *Class) AddStudent(s *Student) {
	for _, student := range c.Students {
		if student.Name == s.Name {
			fmt.Printf("%s用户已经存在\n", s.Name)
			return
		}
	}
	c.Students = append(c.Students, s)
}

func (c *Class) ListStudents() {
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json序列化失败!!!")
		return
	}
	fmt.Printf("json:%s\n", data)
}

func (c *Class) EditStudent(id int, score map[string]int) {
	for _, student := range c.Students {
		if student.Id == id {
			for key, value := range score {
				student.Score[key] = value
			}
		}
	}
}

func (c *Class) DelStudent(id int) {
	std_id := 0
	for index, student := range c.Students {
		if student.Id == id {
			std_id = index
		}
	}
	fmt.Println(std_id)
	if std_id > 0 {
		c.Students = append(c.Students[:std_id], c.Students[std_id+1:]...)
	} else {
		fmt.Println("用户不存在")
	}
}

func main() {
	c1 := NewClass(01, "小一班")
	fmt.Printf("%#v\n", c1)

	s1 := &Student{
		Id:   1,
		Name: "洪子昊",
		Age:  4,
		Score: map[string]int{
			"语文": 99,
			"数学": 100,
			"英语": 98,
		},
	}

	s2 := &Student{
		Id:   2,
		Name: "洪子桉",
		Age:  1,
		Score: map[string]int{
			"语文": 100,
			"数学": 90,
			"英语": 95,
		},
	}

	c1.AddStudent(s1)
	c1.AddStudent(s2)
	c1.EditStudent(1, map[string]int{"语文": 120})
	c1.EditStudent(2, map[string]int{"数学": 98})
	c1.ListStudents()

	c1.DelStudent(2)
	c1.ListStudents()

	str := `{"id":1,"name":"小一班","students":[{"id":1,"name":"洪李敏","age":4,"score":{"数学":100,"英语":98,"语文":120}}]}`
	c2 := &Class{}
	err := json.Unmarshal([]byte(str), c2)
	if err != nil {
		fmt.Println("json反序列化失败")
	}
	c2.ListStudents()
	fmt.Printf("%#v", c2)

}
