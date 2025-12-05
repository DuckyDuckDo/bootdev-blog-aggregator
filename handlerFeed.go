package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DuckyDuckDo/bootdev-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

// Handler function to print XML response RSS Aggregation
func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	xmlResponse, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}
	fmt.Println(xmlResponse)
	return nil
}

// Handler function to add a feed to the current User and updates Db
func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <feed name> <feed url>", cmd.name)
	}

	// Query the current user from DB based on config
	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	// Build out the new feed parameters
	feedName := cmd.args[0]
	feedUrl := cmd.args[1]
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    currUser.ID,
	}

	// Make a Create Query in the Database
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
