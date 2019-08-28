package main

import "fmt"

const (
	tc int = 3511
	uc     = 3511
)

const (
	y9 = 2016 + iota
	y10
	y11
	y12
)

func main() {
	// Printing a number in Decimal, Binary and Hexadecimal
	x := 3511
	fmt.Printf("Decimal : %d\nBinary : %b\nHexadecimal : %#x\n", x, x, x)
	fmt.Println()

	// Expressions using operators
	a := 10
	b := 15

	fmt.Println(a == b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println()

	// Typed and Untyped constants
	fmt.Println(tc, uc)
	fmt.Printf("%T  %T\n", tc, uc)
	fmt.Println()

	// Bit shifting
	y := 3511
	fmt.Printf("Decimal : %d\nBinary : %b\nHexadecimal : %#x\n", y, y, y)
	z := y << 1
	fmt.Printf("Decimal : %d\nBinary : %b\nHexadecimal : %#x\n", z, z, z)
	fmt.Println()

	// Raw string literal
	s := `This
is
a
"raw"
string literal`
	fmt.Println(s)
	fmt.Println()

	//iota
	fmt.Println(y9, y10, y11, y12)
}
