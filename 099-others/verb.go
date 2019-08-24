package main

import "fmt"

func main() {
	i := 10
	s := "ten"
	b := true

	fmt.Printf("Integer : %d, Type %T\n", i, i)
	fmt.Printf("String : %s, Type %T\n", s, s)
	fmt.Printf("Quoted string : %q, Type %T\n", s, s)
	fmt.Printf("Boolean : %t, Type %T\n", b, b)
	fmt.Printf("Any type : %v - type %T : %v - type %T\n", i, i, s, s)
}
