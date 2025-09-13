package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostName(t *testing.T) {
	assert.Equal(t, host, "http://localhost:6969/")
}
