package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
	fmt.Println()

	y := []int{6, 7, 8, 9, 0}

	for i, v := range y {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", y)
	fmt.Println()

	z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(z)
	z = append(z, 52)
	fmt.Println(z)
	z = append(z, 53, 54, 55)
	fmt.Println(z)
	fmt.Println()

	a := []int{56, 57, 58, 59, 60}
	z = append(z, a...)
	fmt.Println(z)
	fmt.Println()

	fmt.Println(z[:5])
	fmt.Println(z[5:])
	fmt.Println()

	z = append(z[:3], z[6:]...)
	fmt.Println(z)
	fmt.Println()

	usStates := make([]string, 50)
	fmt.Println(len(usStates))
	fmt.Println(cap(usStates))
	fmt.Println()

	usStates = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	for i := 0; i < len(usStates); i++ {
		fmt.Println(i, usStates[i])
	}
	fmt.Println()

	mds := [][]string{{"James", "Bond"}, {"Miss", "Hopper"}}
	fmt.Println(mds)

	for _, v := range mds {
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}
	fmt.Println()

	person := map[string][]string{"Allan": {"Golang", "Gaming"}, "James": {"Spying", "Fighting"}}
	fmt.Println(person)

	for k, v := range person {
		fmt.Println(k, v)
	}
	fmt.Println()

	person["Link"] = []string{"Zelda", "Hyrule"}
	fmt.Println(person)
	fmt.Println()

	delete(person, "Link")
	fmt.Println(person)
}
