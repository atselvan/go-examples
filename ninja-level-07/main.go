package main

import "fmt"

func main() {
	x := 5
	fmt.Println(&x)

	fmt.Println()

	p1 := person{
		fn: "Ondarre",
		ln: "Rioja",
	}
	fmt.Println(p1)
	fmt.Println(changeMe(&p1))
}

type person struct {
	fn string
	ln string
}

func changeMe(p *person) person {
	p.fn = "James"
	p.ln = "Bond"
	return *p
}
