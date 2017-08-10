package main

import (
	"fmt"
)

func main() {

	//anonymous struct literal
	tData := struct {
		Jumper string
		Lazy   string
	}{
		Jumper: "elephant",
		Lazy:   "monkey",
	}

	fmt.Printf("%+v\n", tData)
}
