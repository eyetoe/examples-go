package main

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

var blue = color.New(color.FgBlue).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var grey = color.New(color.FgBlack).SprintFunc()

func report(s ...interface{}) {
	fmt.Printf(grey("<--: "))
	fmt.Println(blue(s...))
}

func warn(s ...interface{}) {
	fmt.Printf(grey("<--: "))
	fmt.Println(red(s...))
}

func joy(s ...interface{}) {
	fmt.Printf(grey("<--: "))
	fmt.Println(green(s...))
}

func main() {
	variable := blue("this is from a variable")
	joy("the joy() func take an interface to fmt.Println(), but makes it green")
	report("the report() func takes an interface to fmt.Println(), and makes it blue")
	warn("Syntax for warn() is just like ", "syntax for fmt.Println() ", variable)
	joy(flibbity(323))
	warn(flibbity(7097))
}

func flibbity(y int) string {
	x := 5
	z := x * y
	//report(strconv.Itoa(z))
	return strconv.Itoa(z)
}
