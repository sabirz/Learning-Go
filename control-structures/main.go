package main

import "fmt"

func main() {

	for i := 0; i < 30; i++ {

		if i < 10 {
			fmt.Print(i)
			if i == 9 {
				fmt.Println("")
			}
		} else if i < 20 {
			fmt.Print(i)
			if i == 19 {
				fmt.Println("")
			}

		} else if i < 30 {
			fmt.Print(i)

		}

	}

}
