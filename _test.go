package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock data for testing
var originalURL = "https://example.com"
var shortURL = "abc123"

// TestShortenURL tests the ShortenURL function
func TestShortenURL(t *testing.T) {
	short := ShortenURL(originalURL)
	assert.NotEmpty(t, short, "Shortened URL should not be empty")
}

// TestResolveURL tests the ResolveURL function
func TestResolveURL(t *testing.T) {
	urlMap[shortURL] = originalURL // Simulate stored short URL
	original := ResolveURL(shortURL)
	assert.Equal(t, originalURL, original, "Resolved URL should match the original URL")
}

// TestShortenAndResolve tests the entire shorten and resolve cycle
func TestShortenAndResolve(t *testing.T) {
	short := ShortenURL(originalURL)
	original := ResolveURL(short)
	assert.Equal(t, originalURL, original, "Resolved URL should match the original URL")
}
