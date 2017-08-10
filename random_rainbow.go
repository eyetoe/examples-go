package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rainbow[rand.Intn(len(rainbow))]("magical rainbow"))

}

//fmt.Println(rainbow[rand.Intn(len(rainbow))]("magical rainbow"))
var rainbow = []func(...interface{}) string{
	green,
	blue,
	red,
	yellow,
	cyan,
	magenta,
}

var green = color.New(color.FgGreen).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var cyan = color.New(color.FgCyan).SprintFunc()
var magenta = color.New(color.FgMagenta).SprintFunc()
