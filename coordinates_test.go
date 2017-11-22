package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Equal_ReturnsTrue_InCaseOfEqualCoordinates(t *testing.T) {
	// Arrange
	c := *NewCoordinates(1, 1)
	other := *NewCoordinates(1, 1)

	// Act
	equal := c.Equal(other)

	// Assert
	assert.True(t, equal)
}

func Test_Equal_ReturnsFalse_InCaseOfNotEqualCoordinates(t *testing.T) {
	// Arrange
	c := *NewCoordinates(1, 1)
	other := *NewCoordinates(2, 2)

	// Act
	equal := c.Equal(other)

	// Assert
	assert.False(t, equal)
}
