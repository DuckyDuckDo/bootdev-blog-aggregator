package main

import (
	"fmt"
)

// Struct that defines a specific command consisting of names and arguments
type command struct {
	name string
	args []string
}

// Struct that defines a registry of available commands mapping the name to a handlerFunc that executes when run
type commands struct {
	availableCommands map[string]func(*state, command) error
}

// runs the stated command within the available command registry
func (c *commands) run(s *state, cmd command) error {
	// Checks for existence of cmd in available commands
	if _, ok := c.availableCommands[cmd.name]; !ok {
		return fmt.Errorf("command %s does not exist in directory, please register it", cmd.name)
	}

	// Runs the commands and handles any errors that occurs
	err := c.availableCommands[cmd.name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

// function to register a new command to the registry
func (c *commands) register(name string, f func(*state, command) error) {
	c.availableCommands[name] = f
}
