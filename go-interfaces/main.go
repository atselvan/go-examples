package main

import "fmt"

type Application struct {
	name string
}

type Component struct {
	name   string
	parent string
}

func (a Application) get() string {
	return a.name
}

func (c Component) get() string {
	return c.parent + "." + c.name
}

// Any type that implements the method get() string is also of the type part
// Hence ant value of type Application or Component is also of the type part
type part interface {
	get() string // if you got this method then your my type :P
}

func list(p part) {
	fmt.Println(p.get())
}

func main() {
	a1 := Application{name: "app1"}
	a2 := Application{name: "app2"}
	c1 := Component{name: "comp1", parent: "app1"}
	c2 := Component{name: "comp2", parent: "app1"}
	c3 := Component{name: "comp1", parent: "app2"}

	list(a1)
	list(a2)
	list(c1)
	list(c2)
	list(c3)
}
