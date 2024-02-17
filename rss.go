package main

import (
	"encoding/xml" // Provides functions to encode and decode XML data.
	"io"           // Provides basic interfaces to I/O primitives.
	"net/http"     // Provides HTTP client and server implementations.
	"time"         // Provides functionality for measuring and displaying time.
)

// RSSFeed represents the structure of an RSS feed, including its metadata and items.
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`       // The title of the RSS feed.
		Link        string    `xml:"link"`        // The URL to the RSS feed.
		Description string    `xml:"description"` // A description of the RSS feed.
		Language    string    `xml:"language"`    // The language the RSS feed is written in.
		Item        []RSSItem `xml:"item"`        // A slice of RSS items/posts.
	} `xml:"channel"` // Maps the RSS channel element.
}

// RSSItem represents a single item (or post) within an RSS feed, including its essential elements.
type RSSItem struct {
	Title       string `xml:"title"`       // The title of the item.
	Link        string `xml:"link"`        // The URL to the full item.
	Description string `xml:"description"` // A description or summary of the item.
	PubDate     string `xml:"pubDate"`     // The publication date of the item.
}

// urlToFeed fetches the RSS feed from the given URL and parses it into an RSSFeed struct.
func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second, // Sets a timeout to prevent hanging requests.
	}

	resp, err := httpClient.Get(url) // Performs an HTTP GET request to the feed URL.
	if err != nil {
		return RSSFeed{}, err // Returns an empty RSSFeed and the encountered error.
	}
	defer resp.Body.Close() // Ensures the response body is closed to free resources.

	dat, err := io.ReadAll(resp.Body) // Reads the entire response body.
	if err != nil {
		return RSSFeed{}, err // Returns an empty RSSFeed and the encountered error.
	}
	rssFeed := RSSFeed{} // Initializes an empty RSSFeed struct.

	err = xml.Unmarshal(dat, &rssFeed) // Parses the XML data into the rssFeed struct.
	if err != nil {
		return RSSFeed{}, err // Returns an empty RSSFeed and the encountered error.
	}

	return rssFeed, nil // Returns the populated RSSFeed struct and nil error on success.
}
