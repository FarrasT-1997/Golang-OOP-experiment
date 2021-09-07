package main

import "fmt"

// parent object

type human struct {
	Name    string
	Age     int
	Passion string
}

func (h human) Profession() string {
	if h.Passion == "" {
		return "beggar"
	}
	return "rich people"
}

// child object

type SocialStudent struct {
	human
}

func (s SocialStudent) Profession() string {
	if s.Passion == "geography" {
		return "environment consultan "
	}
	if s.Passion == "economy" {
		return "economy minister"
	}
	if s.Passion == "history" {
		return "historian"
	}
	return "beggar"
}

// child object

type ScienceStudent struct {
	human
}

func (s ScienceStudent) Profession() string {
	if s.Passion == "Math" {
		return "math teacher"
	}
	if s.Passion == "Computer" {
		return "software engineer"
	}
	if s.Passion == "Science" {
		return "engineering"
	}
	return "beggar"
}

func main() {
	human := human{}
	fmt.Println(human.Profession())

	human.Passion = "Science"
	fmt.Println(human.Profession())

	socialStudent := SocialStudent{}
	socialStudent.Passion = "history"

	scienceStudent := ScienceStudent{}
	scienceStudent.Passion = "Math"

	fmt.Println(socialStudent.Profession())
	fmt.Println(scienceStudent.Profession())
}
