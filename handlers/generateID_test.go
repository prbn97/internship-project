package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	id, err := GenerateID(20)
	assert.NoError(t, err)
	assert.Equal(t, 20, len(id))

	_, err = GenerateID(-25)
	assert.Error(t, err)
}
