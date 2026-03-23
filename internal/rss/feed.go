package rss

import "encoding/xml"

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

func NewFeed(channel Channel) Feed {
	return Feed{
		Version: "2.0",
		Channel: channel,
	}
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`

	// Optional fields
	// Language       string `xml:"language,omitempty"`
	// Copyright      string `xml:"copyright,omitempty"`
	// ManagingEditor string `xml:"managingEditor,omitempty"`
	// WebMaster      string `xml:"webMaster,omitempty"`
	// PubDate        string `xml:"pubDate,omitempty"`
	// LastBuildDate  string `xml:"lastBuildDate,omitempty"`
	// Category       string `xml:"category,omitempty"`
	// Generator      string `xml:"generator,omitempty"`
	// Docs           string `xml:"docs,omitempty"`
	// Cloud          string `xml:"cloud,omitempty"`
	// TTL            int    `xml:"ttl,omitempty"`
	// Image          string `xml:"image,omitempty"`
	// Rating         string `xml:"rating,omitempty"`
	// TextInput      string `xml:"textInput,omitempty"`
	// SkipHours      string `xml:"skipHours,omitempty"`
	// SkipDays       string `xml:"skipDays,omitempty"`

	Items []Item `xml:"item"`
}

type Item struct {
	// Title and Description: At least one of these fields must be provided.
	Title       string `xml:"title"`
	Description string `xml:"description"`

	// Link      string `xml:"link"`
	// Author    string `xml:"author,omitempty"`
	// Category  string `xml:"category,omitempty"`
	// Comments  string `xml:"comments,omitempty"`
	// Enclosure string `xml:"enclosure,omitempty"`
	// Guid      string `xml:"guid,omitempty"`
	// PubDate   string `xml:"pubDate,omitempty"`
	// Source    string `xml:"source,omitempty"`
}
