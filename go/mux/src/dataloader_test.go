package apifarm_test

import (
	mocks "apifarm/mocks/src"
	apifarm "apifarm/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONFileLoaderLoadSuccessful(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockFileUtils := new(mocks.FileUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewJSONFileLoaderWithUtils(mockStorage, mockJSON, mockFileUtils, mockQueryFactory)

	path := "PATH TO DATA FILE"
	data := []byte{12, 8, 29}
	vg1 := apifarm.VideoGame{Name: "Conquer Rome I"}
	vg2 := apifarm.VideoGame{Name: "Conquer Rome II"}
	vg3 := apifarm.VideoGame{Name: "Conquer Rome III"}
	videoGames := []apifarm.VideoGame{vg1, vg2, vg3}
	expectedQuery := apifarm.Query{}

	mockFileUtils.On("Read", path).Return(data, nil)
	mockJSON.On("DeserializeVideoGames", data).Return(&videoGames, nil)
	mockStorage.On("AddVideoGame", vg1)
	mockStorage.On("AddVideoGame", vg2)
	mockStorage.On("AddVideoGame", vg3)
	mockQueryFactory.On("BuildMessage", apifarm.SuccessfullyLoadedData, uint(0)).Return(expectedQuery)

	// Act
	actualQuery := subject.Load(path)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockJSON.AssertExpectations(t)
	mockFileUtils.AssertExpectations(t)
}
