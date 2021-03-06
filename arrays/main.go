package main

import (
	"fmt"
)

func main() {
	var nums [5]int
	fmt.Println("empty:", nums)

	nums[4] = 100
	fmt.Println("set:", nums)
	fmt.Println("get:", nums[4])

	ints := []int{1, 2, 3, 4, 5}
	fmt.Println("new slice:", ints)

	ints = append(ints, 6)
	fmt.Println("appended:", ints)

	fmt.Println("0-2", ints[:2])
	fmt.Println("2-4", ints[2:4])
	fmt.Println("4-6", ints[4:])

	for i, val := range ints {
		fmt.Println(i, val)
	}

	fmt.Println(sum(ints))

	fmt.Println("check:", ints)

}

func sum(ints []int) int {
	sumtotal := ints[0] + ints[1] + ints[2] + ints[3] + ints[4] + ints[5]
	return sumtotal
}
