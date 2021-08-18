package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) sayHi() {
	fmt.Println("Hi, i am a human. ")
}

func (h *Human) sayAge() {
	fmt.Println("My age is ", h.age)
}

func (s *Student) sayHi() {
	fmt.Println("Hi, my name is ", s.name, "my school is ", s.school)
}

type Men interface {
	sayHi()
	sayAge()
}

func main() {
	s1 := Student{Human{"xfy", 18}, "Kali"}
	s1.sayHi()
	s1.sayAge()

	var i Men
	i = &s1
	i.sayHi()
	i.sayAge()
}
