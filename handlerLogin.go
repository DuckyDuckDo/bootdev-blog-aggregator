package main

import (
	"context"
	"fmt"
)

// Handles the Logging In command ensuring correct arguments and error handling
func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	name := cmd.args[0]
	user, err := s.db.GetUser(context.Background(), name)
	// If user does not exist in database, it generates an error
	if err != nil {
		return err
	}

	// Updates the configuration file with the new logged in user
	err = s.cfg.SetUser(user.Username)
	if err != nil {
		return err
	}

	fmt.Printf("%s logged in successfully", s.cfg.CurrentUserName)
	return nil
}
