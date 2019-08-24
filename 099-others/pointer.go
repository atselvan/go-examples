package main

import "fmt"

func main() {
	a := 10 // variable declaration
	p := &a // "p (type *int) points to a (type int)" or "p contains the address of a (&a)"
	fmt.Printf("Address of a : %v\n", p)
	fmt.Printf("Value of a : %v\n", *p)

	fmt.Println(&a == p, a == *p) // test

	// the zero value of a pointer is null

	// two pointers are equal only if they point to the same variable or both are nil

	var x, y int
	px := &x // "px points to the address of x"
	py := &y // "py points to the address of y"

	// test
	fmt.Println(px == px)
	fmt.Println(px == py)

	// It is perfectly safe for a function to return the address of a local variable

	var pf = f()
	fmt.Println(pf)

	// each call of the function f returns a distinct value

	fmt.Println(f(), f(), f())

	// Because a pointer contains the address of a variable, passing a pointer argument to a function makes it
	// possible for the function to update the variable that was indirectly passed.

	v := 1
	fmt.Println(incr(&v), v, incr(&v), v)

	// Pointer aliasing is useful because it allows us to access a variable without using its name, but this is a
	// double edge sword: to find all the statements that access a variable, we have to know all its aliases.
}

func f() *int{
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}