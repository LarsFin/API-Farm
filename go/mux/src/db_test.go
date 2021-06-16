package apifarm_test

import (
	apifarm "apifarm/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryGetAllVideoGames(t *testing.T) {
	// Arrange
	expected := []apifarm.VideoGame{{Name: "Lady's Quest 0"}}
	subject := apifarm.NewInMemoryWithVideoGames(&expected)

	// Act
	got := subject.GetAllVideoGames()

	// Assert
	assert.Equal(t, expected, got, "they should be equal")
}

func TestInMemoryAddVideoGame(t *testing.T) {
	// Arrange
	videoGames := []apifarm.VideoGame{}
	subject := apifarm.NewInMemoryWithVideoGames(&videoGames)
	expected := apifarm.VideoGame{Name: "The Great Gamesby"}

	// Act
	got := subject.AddVideoGame(expected)

	// Assert
	assert.Equal(t, expected, got)
	assert.Contains(t, videoGames, expected)
}
