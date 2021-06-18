package apifarm_test

import (
	mocks "apifarm/mocks/src"
	apifarm "apifarm/src"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const path = "PATH TO DATA FILE"

func TestJSONFileLoaderLoadSuccessful(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockFileUtils := new(mocks.FileUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewJSONFileLoaderWithUtils(mockStorage, mockJSON, mockFileUtils, mockQueryFactory)

	data := []byte{12, 8, 29}
	vg1 := apifarm.VideoGame{Name: "Conquer Rome I"}
	vg2 := apifarm.VideoGame{Name: "Conquer Rome II"}
	vg3 := apifarm.VideoGame{Name: "Conquer Rome III"}
	videoGames := []apifarm.VideoGame{vg1, vg2, vg3}
	expectedQuery := apifarm.Query{}

	mockFileUtils.On("Read", path).Return(data, nil)
	mockJSON.On("DeserializeVideoGames", data).Return(&videoGames, nil)
	mockStorage.On("AddVideoGame", vg1).Return(apifarm.VideoGame{})
	mockStorage.On("AddVideoGame", vg2).Return(apifarm.VideoGame{})
	mockStorage.On("AddVideoGame", vg3).Return(apifarm.VideoGame{})
	mockQueryFactory.On("BuildMessage", apifarm.SuccessfullyLoadedData, uint(0)).Return(expectedQuery)

	// Act
	actualQuery := subject.Load(path)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockFileUtils.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockStorage.AssertExpectations(t)
}

func TestJSONFileLoaderLoadFileReadFailure(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockFileUtils := new(mocks.FileUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewJSONFileLoaderWithUtils(mockStorage, mockJSON, mockFileUtils, mockQueryFactory)

	err := errors.New("failed to read file")
	expectedQuery := apifarm.Query{}

	mockFileUtils.On("Read", path).Return(nil, err)
	mockQueryFactory.On("Error", err).Return(expectedQuery)

	// Act
	actualQuery := subject.Load(path)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockFileUtils.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockStorage.AssertExpectations(t)
}

func TestJSONFileLoaderLoadDeserializeFailure(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockFileUtils := new(mocks.FileUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewJSONFileLoaderWithUtils(mockStorage, mockJSON, mockFileUtils, mockQueryFactory)

	data := []byte{12, 9, 34}
	err := errors.New("failed to deserialize")
	expectedQuery := apifarm.Query{}

	mockFileUtils.On("Read", path).Return(data, nil)
	mockJSON.On("DeserializeVideoGames", data).Return(nil, err)
	mockQueryFactory.On("Error", err).Return(expectedQuery)

	// Act
	actualQuery := subject.Load(path)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockFileUtils.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockStorage.AssertExpectations(t)
}
