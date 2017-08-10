package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Search func(player_id string) Result

type Result string

var (
	Bull1 = pusher()
	Bull2 = pusher()
	Bull3 = pusher()
)

func pusher() Search {
	return func(player_id string) Result {

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		return Result(fmt.Sprintf("Result for %q\n", player_id))
	}
}

func main() {
	all := []string{"123456", "234567", "345678", "456789", "567890"}

	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	//var results []Result

	for i := 0; i < len(all); i++ {
		fmt.Println(PushIt(all[i]))
	}

	elapsed := time.Since(start)
	//fmt.Println(results[0])
	fmt.Println(elapsed)
}

func PushIt(player_id string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Bull1(player_id + "-1") }()
	go func() { c <- Bull2(player_id + "-2") }()
	go func() { c <- Bull3(player_id + "-3") }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func fanOut(player_id string) (results []Result) {
	c1 := make(chan Result)
	c2 := make(chan Result)
	c3 := make(chan Result)

	select {
	case v1 := <-c1:
		fmt.Println("using channel 1:", v1)
	case v2 := <-c2:
		fmt.Println("using channel 2:", v2)
	case v3 := <-c3:
		fmt.Println("using channel 3:", v3)
	default:
		fmt.Println("no ready channels.")

	}

	return
}
