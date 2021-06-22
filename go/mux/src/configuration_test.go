package apifarm_test

import (
	mocks "apifarm/mocks/src"
	apifarm "apifarm/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	// Arrange
	mockJSON := new(mocks.DataUtils)
	mockFileUtils := new(mocks.FileUtils)

	path := "PATH TO CONFIG FILE"
	configData := []byte{24, 66, 80}
	expected := apifarm.Configuration{Host: "localhost", Port: 2550}

	mockFileUtils.On("Read", path).Return(configData, nil)
	mockJSON.On("DeserializeConfiguration", configData).Return(&expected, nil)

	// Act
	got, err := apifarm.GetConfiguration(path, mockJSON, mockFileUtils)

	// Assert
	assert.Equal(t, expected, got)
	assert.Nil(t, err)
}
