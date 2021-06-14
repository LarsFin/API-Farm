package apifarm_test

import (
	apifarm "apifarm/src"
	"reflect"
	"testing"
)

func TestInMemoryGetAllVideoGames(t *testing.T) {
	// Arrange
	expected := []apifarm.VideoGame{{Name: "Lady's Quest 0"}}
	subject := apifarm.NewInMemoryWithVideoGames(&expected)

	// Act
	got := subject.GetAllVideoGames()

	// Assert
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("InMemory::GetAllVideoGames failed, expected %v, got %v", expected, got)
	}
}
