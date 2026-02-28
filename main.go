package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}
		cmd, ok := commands[userInput[0]]
		if !ok {
			fmt.Print("Unknown command\n\n")
			continue
		}
		err := cmd.callback()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println()
	}
}
