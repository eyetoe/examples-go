package main

import "fmt"

type Puppy int
type Dingo Puppy

func main() {

	fmt.Println("howdy")

	var dog Dingo = 3

	fmt.Println("howdy", dog)
	fmt.Printf("dog is type: %T\n", dog)
}
