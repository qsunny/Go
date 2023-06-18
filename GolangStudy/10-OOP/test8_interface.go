package main

import "fmt"

type CountryA struct {
	Name string
}
type CityB struct {
	Name string
}
type Stringable interface {
	ToString() string
}

func (c CountryA) ToString() string {
	return "Country = " + c.Name
}
func (c CityB) ToString() string {
	return "City = " + c.Name
}
func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}

func main() {
	d1 := CountryA{"USA"}
	d2 := CityB{"Los Angeles"}
	PrintStr(d1)
	PrintStr(d2)
}
