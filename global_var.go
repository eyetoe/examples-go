package main

import "fmt"

var thingy string = "howdy"

func thingLooker() {

	fmt.Println("I also see thingy as:", thingy)
}

func main() {
	fmt.Println("I see thingy as:", thingy)
	thingLooker()
}
