package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		inputWords := cleanInput(scanner.Text())
		if len(inputWords) == 0 {
			continue
		}
		fmt.Printf("Your command was: %v", inputWords[0])
		fmt.Print("\n")
	}
}
