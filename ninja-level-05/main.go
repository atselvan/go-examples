package main

import "fmt"

type person struct {
	fn                 string
	ln                 string
	age                int
	favIceCreamFlavour []string
}

func main() {
	p1 := person{
		fn:                 "Geralt",
		ln:                 "of Rivia",
		age:                35,
		favIceCreamFlavour: []string{"Chocolate", "Vanilla"},
	}

	p2 := person{
		fn:                 "Yennefer",
		ln:                 "of Vengerberg",
		age:                10,
		favIceCreamFlavour: []string{"Strawberry", "Vanilla"},
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.fn, p1.ln, p1.age, p1.favIceCreamFlavour)
	fmt.Println(p2.fn, p2.ln, p2.age, p2.favIceCreamFlavour)

	fmt.Println()

	m := map[string]person{
		p1.fn: p1,
		p2.fn: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println()

	type vehicle struct {
		color string
		doors int
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t := truck{
		vehicle: vehicle{
			color: "Red",
			doors: 2,
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			color: "Blue",
			doors: 4,
		},
		luxury: true,
	}

	fmt.Println(t, s)
	fmt.Println(t.color, t.doors, t.fourWheel)
	fmt.Println(s.color, s.doors, s.luxury)
}
