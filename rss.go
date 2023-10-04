package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

// Rss type
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

// Rss item type
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// Function to convert a URL to an RSS feed
func urlToFeed(url string) (RSSFeed, error) {
	// Create a new HTTP client with a 10-second timeout.
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	// Make a GET request to the URL.
	resp, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	// Close the response body.
	defer resp.Body.Close()
	//Read the response body into a byte slice.
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, err
	}
	// Create a new RSS feed.
	rssFeed := RSSFeed{}
	// Unmarshal the XML data into the RSS feed.
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}
	// Return the RSS feed.
	return RSSFeed{}, nil
}
