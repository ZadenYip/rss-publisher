package rss

import (
	"fmt"
	"os"

	"zadenyip.com/rss-publisher/pkg/common"
	"zadenyip.com/rss-publisher/pkg/plugin"
)

func GenerateFeed(outputPath string) {
	config := readToml()
	feed := buildFeed(config)

	feed.Channel.Items = buildItems()

	// TODO write feed to output file or stdout
}

func readToml() *Config {
	config, err := ReadChannelConfig("./config.toml")
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
	}

	return config
}

// Impletement in RSS Specification 2.0
func buildFeed(config *Config) common.Feed {
	return common.Feed{
		Version: "2.0",
		Channel: buildChannel(
			config.Channel.Required.Title,
			config.Channel.Required.Link,
			config.Channel.Required.Description,
		),
	}
}

// with empty items list
func buildChannel(title string, link string, description string) *common.Channel {
	return &common.Channel{
		Title:       title,
		Link:        link,
		Description: description,
		Items:       nil,
	}
}

// with check
func buildItems() []common.Item {
	plg := plugin.GetPlugin()
	items := plg.BuildItems()

	for i, item := range items {
		if err := validateItem(item); err != nil {
			fmt.Printf("Invalid item %d: %v\n", i, err)
			os.Exit(1)
		}
	}

	return items
}

func validateItem(item common.Item) error {
	if item.Title == "" && item.Description == "" {
		return fmt.Errorf("item must have at least a title or a description")
	}

	return nil
}
