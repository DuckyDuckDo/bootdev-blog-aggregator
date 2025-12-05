package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DuckyDuckDo/bootdev-blog-aggregator/internal/database"
	"github.com/google/uuid"
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

// Handler function that ends up listing all users
func handlerUsers(s *state, cmd command) error {
	// Ensures correct usage of the command
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	// Query the database
	allUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	// Loop through all users and print their username, checking for current logged in user
	for _, user := range allUsers {
		if user.Username == s.cfg.CurrentUserName {
			fmt.Printf("- %s (current)\n", user.Username)
		} else {
			fmt.Printf("- %s\n", user.Username)
		}
	}
	return nil
}
