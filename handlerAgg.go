package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	xmlResponse, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}
	fmt.Println(xmlResponse)
	return nil
}
