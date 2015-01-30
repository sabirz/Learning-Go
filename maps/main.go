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

	fmt.Println("")
	fmt.Println(keys(m))
}

func keys(m map[string]int) ([]string, []int) {

	keys := []string{}
	v := []int{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	for _, l := range m {
		v = append(v, l)
	}

	return keys, v

}
