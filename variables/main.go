package main

import "fmt"

var (
	author  = "@sabirz"
	Version = "0.0.1"
)

const (
	CCVisa       = "Visa"
	CCMasterCard = "MasterCard"
	CCAmex       = "American Express"
)

func main() {

	var greeting string = "Hello World!"
	var a, b, c int = 1, 2, 3

	fmt.Println(greeting, a, b, c)

	var name = "Sabir"
	var d, e, f = 1, 2.0, "Three"

	fmt.Println(name, d, e, f)

	Course := "Learning Go"
	x, y, z := 5, 6, 7

	fmt.Println(Course, x, y, z)
}
