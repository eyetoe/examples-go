package main

import (
	"bytes"
	"fmt"
	"html/template"
)

func main() {

	message := "The quick brown {{.Jumper}} jumps over the lazy {{.Lazy}}.\n"

	//anonymous struct literal
	tData := struct {
		Jumper string
		Lazy   string
	}{
		Jumper: "elephant",
		Lazy:   "monkey",
	}

	fmt.Printf("%+v\n", tData)

	tmpl, err := template.New("message").Parse(message)
	if err != nil {
		panic(err)
	}
	var tOut bytes.Buffer

	if err := tmpl.Execute(&tOut, tData); err != nil {
		panic(err)
	}
	fmt.Printf(tOut.String())
}
