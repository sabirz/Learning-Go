package main

import (
	"fmt"
)

func main() {
	x, y := 20, 40

	num := 5
	double(num)
	fmt.Println(num)

	triple(&num)
	fmt.Println(num)

	swap(&x, &y)
	fmt.Println(x, y)

}

func double(n int) {
	n = n * 2
}

func triple(n *int) {
	*n = *n * 3
}

func swap(x, y *int) {
	*x = *x * 5
	*y = *y * 5
}
