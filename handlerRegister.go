package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DuckyDuckDo/bootdev-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

// Handles the Logging In command ensuring correct arguments and error handling
func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	// Grab the name from user input, and build out the parameters for inserting user into DB
	name := cmd.args[0]
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  name,
	}

	// Calls the query after sqlc generates the code for the CRUD application
	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}

	// Sets the name in the config file after adding to database
	err = s.cfg.SetUser(user.Username)
	if err != nil {
		return err
	}

	fmt.Printf("%s registered in successfully", user.Username)
	return nil
}
