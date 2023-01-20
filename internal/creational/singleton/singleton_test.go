package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	got := New()
	assert.NotNil(t, got)
}
