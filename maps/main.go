package main

import (
	"fmt"
)

func main() {

	age := make(map[string]int)

	age["jeremy"] = 24
	age["jordie"] = 21
	age["josh"] = 27

	fmt.Println(age)

	fmt.Println("Jeremy's age:", age["jeremy"])

	delete(age, "jeremy")
	fmt.Println(age)

	m := map[string]int{
		"jeremy": 24,
		"jordie": 22,
		"josh":   27,
	}
	fmt.Println(m)

	for n, a := range m {
		fmt.Printf("%v is %v years old", n, a)

	}

	fmt.Println(keys(m))
}

func keys(s string, i int) {

	k := []string{age["jeremy"], age["jordie"], age["josh"]}

	return k

}
