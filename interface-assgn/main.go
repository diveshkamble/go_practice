package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (s square) getArea() float64 {
	return (s.sideLength * s.sideLength)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func main() {

	sq1 := square{sideLength: 5.0}
	printArea(sq1)
	tr1 := triangle{base: 5.0, height: 10.0}
	printArea(tr1)
}
