package rss

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Channel struct {
		Required struct {
			Title       string
			Link        string
			Description string
		} `toml:"required"`

		// Optional struct {

		// } `toml:"optional"`
	} `toml:"channel"`
}

func ReadChannelConfig(path string) (*Config, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return &config, nil
}
