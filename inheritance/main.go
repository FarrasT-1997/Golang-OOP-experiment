package main

import "fmt"

type human struct {
	Name    string
	Age     int
	Address string
}

type student struct {
	human
	Class string
}

type teacher struct {
	human
	skill string
}

func main() {
	student1 := student{
		Class: "1A",
	}
	student1.Age = 17
	student1.Name = "farras"
	student1.Address = "Depok"

	fmt.Println(student1)

	teacher1 := teacher{
		skill: "Math",
	}
	teacher1.Age = 30
	teacher1.Name = "muti"
	teacher1.Address = "Sukabumi"

	fmt.Println(teacher1)
}
