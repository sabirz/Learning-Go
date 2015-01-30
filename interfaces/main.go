package main

import (
	"fmt"
)

type Cat struct {
	name string
}

func (c Cat) Pet() {
	fmt.Println("prrrrr")
}

func (c Cat) Name() string {
	return c.name
}

type Dog struct {
	name string
}

func (d Dog) Pet() {
	fmt.Println("woof woof")
}

func (d Dog) Name() string {
	return d.name
}

type Sheep struct {
	name string
}

func (s Sheep) Pet() {
	fmt.Println("baah baah")
}

func (s Sheep) Name() string {
	return s.name
}

type Animal interface {
	Pet()
	Name() string
}

func Compliment(a Animal) {
	fmt.Println("Great Job", a.Name())
	a.Pet()
}

func main() {

	c := Cat{"johnny larry"}
	Compliment(c)

	d := Dog{"trixie belle"}
	Compliment(d)

	s := Sheep{"Dolly Clony"}
	Compliment(s)

}
