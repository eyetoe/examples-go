package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "Izro.json"
	fmt.Println(strings.Replace(text, ".json", "", -1))
	// prints: Izro

}
