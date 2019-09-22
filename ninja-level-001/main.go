package main

import "fmt"

type myType int

var x myType
var y int

func main() {
	// If a VALUE for the VARIABLE is not defined post DECLARATION
	// a ZERO VALUE will be ASSIGNED to each variable
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
