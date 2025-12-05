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
	// Checks for correct command usage
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	// Calls fetchFeed function to grab XML response
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

func handlerFeeds(s *state, cmd command) error {
	// Checks for proper usage
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	// Queries the Feeds Table for all Feeds
	allFeeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range allFeeds {
		// Prints feed name and URL
		fmt.Printf("- %s \n", feed.Name)
		fmt.Printf("- %s \n", feed.Url)

		// Fetches the user based on ID and prints user
		feedUser, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Printf("- %s \n", feedUser.Username)

	}

	return nil
}
