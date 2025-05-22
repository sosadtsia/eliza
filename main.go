package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sosadtsia/eliza/doctor"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	intro := doctor.Intro()

	fmt.Println(intro)

	fmt.Println("You can type 'quit' to exit the program.")
	fmt.Println("You can type 'help' to see the list of available commands.")
	fmt.Println("You can type 'clear' to clear the screen.")

	for {
		fmt.Print("-> ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		userInput = strings.TrimSpace(userInput)

		switch strings.ToLower(userInput) {
		case "quit":
			fmt.Println("Goodbye!")
			return
		case "help":
			fmt.Println("Available commands: quit, help, clear")
			continue
		case "clear":
			fmt.Print("\033[H\033[2J")
			fmt.Println("Screen cleared.")
			continue
		}

		response := doctor.Response(userInput)

		if response == "" {
			fmt.Println("I don't understand.")
			continue
		}

		fmt.Println(response)
	}
}
