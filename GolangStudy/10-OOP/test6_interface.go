package main

import "fmt"

type Person struct {
	Name   string
	Sexual string
	Age    int32
}

func PrintPerson(p *Person) {
	fmt.Printf("Name=%s, Sexual=%s,Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

func (p *Person) Print() {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

func main() {
	var p = Person{
		Name:   "Hao Chen",
		Sexual: "Male",
		Age:    44,
	}
	PrintPerson(&p)
	p.Print()
}
