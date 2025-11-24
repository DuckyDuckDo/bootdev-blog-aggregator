package main

import (
	"fmt"
)

// Handles the Logging In command ensuring correct arguments and error handling
func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.name)
	}

	name := cmd.args[0]
	err := s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("%s logged in successfully", s.cfg.CurrentUserName)
	return nil
}
