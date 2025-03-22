package test

import (
	"github.com/Scorpio69t/go-utils/pkg/curl"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewCurl(t *testing.T) {
	// Create a new Curl instance
	c := curl.NewCurl("http://example.com", 10*time.Second)
	assert.NotNil(t, c)
}
