package main

import (
	"testing"
)

func Test_Equal_ReturnsTrue_InCaseOfEqualCoordinates(t *testing.T) {
	// Arrange
	c := *NewCoordinates(1, 1)
	other := *NewCoordinates(1, 1)

	// Act
	equal := c.Equal(other)

	// Assert
	if !equal {
		t.Fatal("coordinates are equal")
	}
}

func Test_Equal_ReturnsFalse_InCaseOfNotEqualCoordinates(t *testing.T) {
	// Arrange
	c := *NewCoordinates(1, 1)
	other := *NewCoordinates(2, 2)

	// Act
	equal := c.Equal(other)

	// Assert
	if equal {
		t.Fatal("coordinates not are equal")
	}
}
