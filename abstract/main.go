package main

import "fmt"

type Hitung interface {
	Area() float32
}

type square struct {
	length float32
}

type circle struct {
	radius float32
}

type rectangle struct {
	length float32
	width  float32
}

func (s square) Area() float32 {
	return s.length * s.length
}

func (c circle) Area() float32 {
	return (3.14 * c.radius * c.radius)
}

func (r rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	var area Hitung

	area = square{length: 10}
	fmt.Println(area.Area())

	area = circle{radius: 5}
	fmt.Println(area.Area())

	area = rectangle{length: 5, width: 10}
	fmt.Println(area.Area())
}
