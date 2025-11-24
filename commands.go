package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	availableCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if _, ok := c.availableCommands[cmd.name]; !ok {
		return fmt.Errorf("command %s does not exist in directory, please register it", cmd.name)
	}

	err := c.availableCommands[cmd.name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.availableCommands[name] = f
}
