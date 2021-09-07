package main

import (
	"farras/controller"
	"fmt"
)

func main() {
	input := controller.Account{
		Name:     "Farras",
		Address:  "Depok",
		Age:      24,
		Password: "1234",
	}
	fmt.Println(input.InputUser())
	// fmt.Println(input.getPassword())  // program cannot share password from this unexported method
}
