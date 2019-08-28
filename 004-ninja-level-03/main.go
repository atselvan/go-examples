package main

import "fmt"

func main() {
	// printing ascii
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t%s\t%#U\n", i, string(i), i)
	}

	fmt.Println()

	//printing numbers from 1 to 100
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	//print every rune code point of the upper case alphabet three times
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

	fmt.Println()

	// for loop - print the years you have been alive
	i := 1993
	for i < 2020 {
		fmt.Println(i)
		i++
	}

	fmt.Println()

	i = 1993
	for {
		if i >= 2020 {
			break
		}
		fmt.Println(i)
		i++
	}

	fmt.Println()

	// print out the remainder (modulus) which is found between 10 and 100
	// when it is divided by 4
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}
