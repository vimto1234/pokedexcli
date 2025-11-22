package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/vimto1234/pokedexcli/internal/pokecache"
)

type config struct {
	next          string
	previous      string
	locationCache pokecache.Cache
}

var mainConfig config = config{
	next:          "https://pokeapi.co/api/v2/location-area/",
	previous:      "",
	locationCache: pokecache.NewCache(5 * time.Second),
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		inputWords := cleanInput(scanner.Text())
		if len(inputWords) == 0 {
			continue
		}

		command := inputWords[0]

		args := []string{}

		if len(inputWords) >= 2 {
			args = inputWords[1:]
		}

		executeCommand(command, args)
		fmt.Print("\n")
	}
}
