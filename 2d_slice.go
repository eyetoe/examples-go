package main

import "fmt"

func main() {
	x := []string{"x0", "x1", "x2", "x3", "x4"}
	y := []string{"y0", "y1", "y2", "y3", "y4"}

	var all [][]string

	all = append(all, x)
	all = append(all, y)

	fmt.Println(x[3])
	fmt.Println(y[3])

	fmt.Println(len(all[0]))
	fmt.Println(all[1][3])

}
