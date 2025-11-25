package main

import (
	"context"
	"fmt"
)

// Handles the Logging In command ensuring correct arguments and error handling
func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	s.db.Reset(context.Background())

	fmt.Println("Resetted the database successfully")
	return nil
}
