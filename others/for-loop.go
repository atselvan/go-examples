package main

import (
	"fmt"
)

// The for loop is the only loop statement in Go, It has a number of forms..

func main() {
	var value int
	a := 1
	s := []int{1, 2, 3, 4, 5}

	fmt.Println("a traditional for loop")
	/*
		 	for initialization; condition; post {
				// zero or more statements
			}
	*/

	for i := 1; i < len(s); i++ {
		fmt.Println(s[i])
	}

	fmt.Println("a traditional while loop")
	/*
		  for condition {
			 // zero or more statements
		 }
	*/

	for value == 4 {
		a += value
		value = a
		fmt.Println(value)
	}
}
