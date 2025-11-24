package main

import (
	"log"
	"os"

	"github.com/DuckyDuckDo/bootdev-blog-aggregator/internal/config"
)

// struct that will maintain app state and can be updated
type state struct {
	cfg *config.Config
}

func main() {
	// Grabs initial config and sets the state with the initial config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("%v", err)
	}

	appState := &state{
		cfg: &cfg,
	}

	// establishes a map of commands
	commandMap := commands{
		availableCommands: make(map[string]func(*state, command) error),
	}

	// Registers a login command
	commandMap.register("login", handlerLogin)

	// checks the user usage of the CLI
	userArgs := os.Args
	if len(userArgs) < 2 {
		log.Fatal("Usage: go run . command args[...]")
		os.Exit(1)
	}

	// parses the user input into a command with various arguments
	cmd := command{
		name: userArgs[1],
		args: userArgs[2:],
	}

	// Executes the command based on user input
	err = commandMap.run(appState, cmd)
	if err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}

}
