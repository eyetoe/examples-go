package main

import "fmt"

func main() {

	x := 10
	y := 22

	if y > 0 && y < 21 && x > 0 && x < 11 {
		fmt.Println("withing range")
	} else {
		fmt.Println("outside range")
	}
}
