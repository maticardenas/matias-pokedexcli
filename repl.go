package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Lists location areas",
			callback:    callbackMap,
		},
		"map-back": {
			name:        "map-back",
			description: "Lists previous location areas",
			callback:    callbackMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Lists Pokemon in a location area",
			callback:    callbackExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exits the application",
			callback:    callbackExit,
		},
	}
}

func startRepl(cfg *config) { // we pass config as a pointer to be able to modify it
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		availableCommands := getCommands()
		cmd, ok := availableCommands[command]

		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := cmd.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
