package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// Uses this function to fetch a feed from a certain URL
func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Set up the HTTP Client
	client := &http.Client{}

	// Set up the request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	// Perform the request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	// Load data for Unmarshalling
	byte_data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response into the RSSFeed Struct format
	var xmlResponse RSSFeed
	if err := xml.Unmarshal(byte_data, &xmlResponse); err != nil {
		return nil, err
	}

	// Clean up the responses
	xmlResponse.Channel.Title = html.UnescapeString(xmlResponse.Channel.Title)
	xmlResponse.Channel.Description = html.UnescapeString(xmlResponse.Channel.Description)
	for i := range xmlResponse.Channel.Item {
		xmlResponse.Channel.Item[i].Title = html.UnescapeString(xmlResponse.Channel.Item[i].Title)
		xmlResponse.Channel.Item[i].Description = html.UnescapeString(xmlResponse.Channel.Item[i].Description)
	}

	return &xmlResponse, nil
}
