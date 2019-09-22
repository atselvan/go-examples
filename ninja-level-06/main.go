package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(foo())
	fmt.Println(bar())

	fmt.Println()

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sumVariadic(numbers...))
	fmt.Println(sumFor(numbers))

	fmt.Println()

	deferInAction()

	fmt.Println()

	p1 := person{
		fn:  "Ondarre",
		ln:  "Rioja",
		age: 50,
	}
	p1.speak()

	fmt.Println()

	sq := square{side: 5}
	c := circle{radius: 2}

	info(sq)
	info(c)

	fmt.Println()
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return mathSquare(s.side)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * mathSquare(c.radius)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("Area :", s.area())
}

func mathSquare(n float64) float64 {
	return n * n
}

type person struct {
	fn  string
	ln  string
	age int
}

func (p person) speak() {
	fmt.Printf("Hi! I am %s %s and I am %d years old\n", p.fn, p.ln, p.age)
}

func deferInAction() {
	defer fmt.Println("Now defer runs")
	fmt.Println("Before defer the function block will execute")
	fmt.Println(".....")
	fmt.Println("....")
	fmt.Println("...")
	fmt.Println("..")
	fmt.Println(".")
}

func sumVariadic(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func sumFor(numbers []int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func foo() int {
	return 5
}

func bar() (int, string) {
	return 5, "favourite number"
}
