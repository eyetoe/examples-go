package main

import (
	"fmt"

	"github.com/fatih/color"
)

var blue = color.New(color.FgBlue).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var wild = color.New(color.BgRed, color.FgYellow).SprintFunc()
var uline = color.New(color.BgBlack, color.FgMagenta, color.Underline).SprintFunc()

func main() {

	fmt.Printf("You can set any %s you like\n", blue("Colors"))
	fmt.Printf("You can set any %s you like\n", green("Colors"))
	fmt.Printf("You can set any %s you like\n", red("Colors"))
	fmt.Printf("You can set any %s you like\n", wild("Colors"))
	fmt.Printf("You can set any %s you like\n", uline("Colors"))

}
