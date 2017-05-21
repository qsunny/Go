package main

import (
	"fmt"
)

type Employee struct {
	Name string
	Age  int
	Rank int
}

//method chain
func (empl *Employee) Promote() *Employee {
	fmt.Printf("Promote %s\n", empl.Name)
	empl.Rank++
	return empl
}

func (empl *Employee) SetName(name string) *Employee {
	fmt.Printf("Set name of new Employee to %s\n", name)
	empl.Name = name
	return empl
}

func main() {
	Bob := new(Employee)
	Bob.Rank = 1
	fmt.Println("Bob rank now is :", Bob.Rank)
	Bob.Promote()
	fmt.Println("Bob rank now is :", Bob.Rank)

	worker := new(Employee)
	worker.Rank = 1
	worker.SetName("Aaron.Qiu").Promote()
	fmt.Printf("Here we have %s with rank %d\n", worker.Name, worker.Rank)

}
