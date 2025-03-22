package test

import (
	"github.com/stretchr/testify/assert"
	"go-utils/pkg/curl"
	"testing"
	"time"
)

func TestNewCurl(t *testing.T) {
	// Create a new Curl instance
	c := curl.NewCurl("http://example.com", 10*time.Second)
	assert.NotNil(t, c)
}
