package main

import "fmt"

type Thingy struct {
	Name  string
	Color string
}

func main() {
	var Banana = Thingy{
		Name: "Banana",
	}

	Banana.Color = "yellow"

	fmt.Printf("%+v\n", Banana)

}
