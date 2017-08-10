package main

import "fmt"

func main() {
Mushroom:
	for x := 0; x < 4; x++ {
		fmt.Println("Mushroom")
	Badger:
		for y := 0; y < 10; y++ {
			if y > 3 {
				break Mushroom
				break Badger
			}
			fmt.Println("badger")
		}
	}

	for z := 0; z < 10; z++ {
		if z > 5 {
			fmt.Println("BAZINGA!")
			// THIS BREAK IS OUT OF CONTEXT
			// cause it's not inside the Mushroom loop
			//break Mushroom
		}
	}
}
