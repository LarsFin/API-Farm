package apifarm_test

import (
	apifarm "apifarm/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryGetVideoGame(t *testing.T) {
	// Arrange
	const id = 5
	subject, videoGames := apifarm.NewInMemoryForTests()
	expected := apifarm.VideoGame{ID: id}
	includingVideoGames := []apifarm.VideoGame{
		{ID: 2},
		expected,
		{ID: 7},
	}
	*videoGames = &includingVideoGames

	// Act
	got := subject.GetVideoGame(id)

	// Assert
	assert.Equal(t, expected, got)
}

func TestInMemoryGetVideoGameNil(t *testing.T) {
	// Arrange
	subject := apifarm.NewInMemory()

	// Act
	got := subject.GetVideoGame(99)

	// Assert
	assert.Nil(t, got)
}

func TestInMemoryGetAllVideoGames(t *testing.T) {
	// Arrange
	subject, videoGames := apifarm.NewInMemoryForTests()
	expected := **videoGames

	// Act
	got := subject.GetAllVideoGames()

	// Assert
	assert.Equal(t, expected, got)
}

func TestInMemoryAddVideoGame(t *testing.T) {
	// Arrange
	subject, videoGames := apifarm.NewInMemoryForTests()
	videoGame := apifarm.VideoGame{Name: "The Great Gamesby"}
	expected := apifarm.VideoGame{ID: uint(1), Name: "The Great Gamesby"}

	// Act
	got := subject.AddVideoGame(videoGame)

	// Assert
	assert.Equal(t, expected, got)
	assert.Contains(t, **videoGames, expected)
}

func TestInMemoryAddVideoGameIncrementsAndSetsId(t *testing.T) {
	// Arrange
	subject, videoGames := apifarm.NewInMemoryForTests()

	// Act
	videoGame1 := subject.AddVideoGame(apifarm.VideoGame{Name: "VIDEO GAME 1"})
	videoGame2 := subject.AddVideoGame(apifarm.VideoGame{Name: "VIDEO GAME 2"})
	videoGame3 := subject.AddVideoGame(apifarm.VideoGame{Name: "VIDEO GAME 3"})

	// Assert
	assert.Equal(t, videoGame1.ID, uint(1))
	assert.Equal(t, videoGame2.ID, uint(2))
	assert.Equal(t, videoGame3.ID, uint(3))

	assert.Len(t, **videoGames, 3)
}
