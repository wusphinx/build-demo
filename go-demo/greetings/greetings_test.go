package greetings_test

import (
	"testing"

	"github.com/go-playground/assert"

	"example_module/greetings"
)

func TestGreeting(t *testing.T) {
	assert.NotEqual(t, greetings.Greeting(), "")
}
