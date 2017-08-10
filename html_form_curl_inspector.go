package main

import (
	"fmt"
	"log"
	"net/http"
)

// poke this with
// curl -X POST --data "name=bob" http://localhost:9090

// handler for requests, GET and POST
func inspect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	r.ParseForm()

	fmt.Println("Full Form:", r.Form)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("name:", r.Form["name"])
}

// main does routing
func main() {
	fmt.Println("running on: http://localhost:9090 for testing")
	http.HandleFunc("/", inspect)            // setting router rule
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
