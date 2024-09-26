package tests

import (
	"shorty/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortener__InvalidUrlGenerateShortURL__ExpectedErrors(t *testing.T) {
	// FIXTURES
	invalidUrl := "invalid_url"
	shortenerService := services.NewShortenerService()

	// EXERCISE
	shortUrl, err := shortenerService.GenerateShortURL(invalidUrl)

	// ASSERTS
	assert.Empty(t, shortUrl)
	assert.EqualError(t, err, "invalid url")
}
