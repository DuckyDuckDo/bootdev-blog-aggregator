package main

import (
	"log"
	"os"

	"github.com/DuckyDuckDo/bootdev-blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("%v", err)
	}

	appState := &state{
		cfg: &cfg,
	}

	commandMap := commands{
		availableCommands: make(map[string]func(*state, command) error),
	}

	userArgs := os.Args
	if len(userArgs) < 2 {
		log.Fatal("Usage: go run . command args[...]")
		os.Exit(1)
	}

	cmd := command{
		name: userArgs[1],
		args: userArgs[2:],
	}

	commandMap.register("login", handlerLogin)
	err = commandMap.run(appState, cmd)
	if err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}

}
