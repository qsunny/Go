package main

import "fmt"

type WithName struct {
	Name string
}

//	type Country struct {
//		Name string
//	}
//
//	type City struct {
//		Name string
//	}

type Country struct {
	WithName
}
type City struct {
	WithName
}

type Printable interface {
	PrintStr()
}

func (w WithName) PrintStr() {
	fmt.Println(w.Name)
}

func (c Country) PrintStr() {
	fmt.Println(c.Name)
}
func (c City) PrintStr() {
	fmt.Println(c.Name)
}

func main() {
	//c1 := Country{"China"}
	//c2 := City{"Beijing"}
	//c1.PrintStr()
	//c2.PrintStr()

	c1 := Country{WithName{"China"}}
	c2 := City{WithName{"Beijing"}}
	c1.PrintStr()
	c2.PrintStr()
}
