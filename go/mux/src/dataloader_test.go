package apifarm_test

import (
	mocks "apifarm/mocks/src"
	apifarm "apifarm/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONFileLoaderLoadSuccessful(t *testing.T) {
	// Arrange
	mockJSON := new(mocks.DataUtils)
	mockFileUtils := new(mocks.FileUtils)

	subject := apifarm.NewJSONFileLoaderWithUtils(mockJSON, mockFileUtils)

	path := "PATH TO DATA FILE"
	data := []byte{12, 8, 29}
	expected := []apifarm.VideoGame{}

	mockFileUtils.On("Read", path).Return(data, nil)
	mockJSON.On("DeserializeVideoGames", data).Return(expected, nil)

	// Act
	actual := subject.Load(path)

	// Assert
	assert.Equal(t, expected, actual)
	mockJSON.AssertExpectations(t)
	mockFileUtils.AssertExpectations(t)
}
