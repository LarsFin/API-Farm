package apifarm_test

import (
	apifarm "apifarm/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryGetAllVideoGames(t *testing.T) {
	// Arrange
	subject, videoGames := apifarm.NewInMemoryForTests()
	expected := **videoGames

	// Act
	got := subject.GetAllVideoGames()

	// Assert
	assert.Equal(t, expected, got, "they should be equal")
}

func TestInMemoryAddVideoGame(t *testing.T) {
	// Arrange
	subject, videoGames := apifarm.NewInMemoryForTests()
	expected := apifarm.VideoGame{Name: "The Great Gamesby"}

	// Act
	got := subject.AddVideoGame(expected)

	// Assert
	assert.Equal(t, expected, got)
	assert.Contains(t, **videoGames, expected)
}
