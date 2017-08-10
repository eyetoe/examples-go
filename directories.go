package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//var FBB_HOME string

var HOME, SAVES = FbbDirs()

func main() {
	fmt.Println("I see foobarbaz home =", HOME)
	fmt.Println("I see save dir =", SAVES)
	EnvSetup()
	return
}

// return the HOME and SAVE vars values
func FbbDirs() (string, string) {
	// Is the ENV var $HOME set?
	home, ok := os.LookupEnv("HOME")
	if ok != true {
		fmt.Println("Please set your $HOME environment variable.")
		os.Exit(1)
	}
	return home + "/.footest/", home + "/.footest/saves/"
}

// Prepare the game HOME directory
func EnvSetup() {
	// Is the client dir already there?
	if _, err := os.Stat(HOME); err != nil {
		if os.IsNotExist(err) {
			// Make the dir
			if err := os.Mkdir(HOME, 0744); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	// Is the saves dir already there?
	if _, err := os.Stat(SAVES); err != nil {
		if os.IsNotExist(err) {
			// Make the dir
			if err := os.Mkdir(SAVES, 0744); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
	// set working dir to $HOME/.foobarbaz
	if err := os.Chdir(HOME); err != nil {
		fmt.Println("Can't move to working directory", err)
	}

	// TESTING: list files in working dir
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fmt.Println("list files:", f.Name())
	}
	return
}
