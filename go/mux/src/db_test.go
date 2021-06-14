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
