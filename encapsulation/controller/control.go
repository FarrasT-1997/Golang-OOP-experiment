package controller

import "fmt"

type Account struct {
	Name     string
	Address  string
	Age      int
	Password string // encapsulated variable
}

func (a Account) InputUser() string {
	output := fmt.Sprintf("%s lives in %s and he/she is %d years old. His/her password is %s", a.Name, a.Address, a.Age, a.getPassword())
	return output
}

func (a Account) getPassword() string {
	return a.Password
}
