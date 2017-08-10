package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//	quest := "What's your name"
	answer := PromptConfirm("What's your name? ")
	fmt.Println("Your name is:", answer)
}

//PromptConfirm returns a user entered string with confirmation from
// the user.  e.g.
//answer := PromptConfirm("What's your name? ")
//fmt.Println("Your name is:", answer)
func PromptConfirm(question string) string {
	// assign var input outside of loop so that it's in scope for return
	var response string
Ask:
	for {
		fmt.Printf("%s", question)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		response = scanner.Text()

		// Remove all non alpha characters, including spaces
		response = stripchars(response, " 1234567890,>?<|/{}[]=+-_*&^%$#@!/(/)\\")
		// don't allow empty
		if response == "" {
			continue
		}

	Confirm:
		for {
			fmt.Printf("You choose: %s. confirm y/n? > ", response)

			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			choice := string([]byte(input)[0])

			switch choice {
			case "y", "Y":
				break Ask
			case "n", "N":
				break Confirm
			}
		}
	}
	return response
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}
