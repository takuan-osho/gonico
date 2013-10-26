package gonico

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	assert.Equal(t, Login(), true)
}
