package main

import "fmt"

type matrix [][]int

//const size int = 4

func main() {
	a := matrix{{1, 2, 3, 4}, {11, 12, 13, 14}}

	fmt.Println("a is:", a)
	fmt.Println()

	a = changer(a)
	fmt.Println("a is now:", a)
	fmt.Println(len(a))
	fmt.Println(len(a[0]))

	for t := range a[0] {
		fmt.Println("t is", t)
	}
	return
}

//func changer(a matrix) matrix {
func changer(a matrix) matrix {
	for xi, x := range a {
		fmt.Println("xi, i:", xi, x)
		for yi, y := range x {
			fmt.Println("yi, y:	", yi, y)
			a[xi][yi] = y * 10
		}
	}
	return a
}

func build() matrix {
	var d matrix
	return d
}
