package rss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	const path = "./fixtures/config.toml"
	config, err := ReadChannelConfig(path)
	assert.NoError(t, err)

	assert.Equal(t, "Channel Title", config.Channel.Required.Title)
	assert.Equal(t, "https://www.github.com", config.Channel.Required.Link)
	assert.Equal(t, "A test RSS channel", config.Channel.Required.Description)
}

func TestReadNonExistentConfig(t *testing.T) {
	const path = "./fixtures/nonexistent.toml"
	_, err := ReadChannelConfig(path)

	assert.ErrorContains(t, err, "failed to read config file")
}

func TestReadInvalidConfig(t *testing.T) {
	const path = "./fixtures/invalid_config.toml"
	_, err := ReadChannelConfig(path)

	assert.ErrorContains(t, err, "failed to unmarshal config file:")
}
