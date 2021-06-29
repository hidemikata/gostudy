package main

import (
	"fmt"
)

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	fmt.Println("hello")
	return p.Name
}
func DriveCar(h Human) {
	if h.Say() == "mikata" {
		fmt.Println("ok")
	} else {
		fmt.Println("ng")
	}
}
func main() {
	var h Human = &Person{"mikata"}
	h.Say()
	DriveCar(h)
	var m Human = &Person{"mikata2"}
	DriveCar(m)

}
