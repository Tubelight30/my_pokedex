package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/chzyer/readline"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, args ...string) error
}

func startRepl(c *config) {
	// scanner := bufio.NewScanner(os.Stdin)
	config := &readline.Config{
		Prompt:            "Pokedex > ",
		HistoryFile:       "/home/saumyagupta/my_pokedex/history.tmp", // Temporary file to store history
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
	}

	rl, err := readline.NewEx(config)
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()
	for {
		// fmt.Print("Pokedex > ")
		// scanner.Scan()
		// text := scanner.Text()
		line, err := rl.Readline()
		if err != nil { // Handle errors such as Ctrl+C or EOF
			log.Fatal(err)
		}
		words := cleanInput(line)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(c, args...)

			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
			continue
		}

	}

}
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "list of all the Pokémon in a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokeom",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Learn stats of a pokemon you already caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See all the pokemon you have caught",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	output = strings.TrimSpace(output)
	words := strings.Fields(output)
	return words
}
