package apifarm_test

import (
	mocks "apifarm/mocks/src"
	apifarm "apifarm/src"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const configPath = "PATH TO CONFIG FILE"

func TestGetConfiguration(t *testing.T) {
	// Arrange
	mockJSON := new(mocks.DataUtils)
	mockFileUtils := new(mocks.FileUtils)

	configData := []byte{24, 66, 80}
	expected := apifarm.Configuration{Host: "localhost", Port: 2550}

	mockFileUtils.On("Read", configPath).Return(configData, nil)
	mockJSON.On("DeserializeConfiguration", configData).Return(&expected, nil)

	// Act
	got, err := apifarm.GetConfigurationForTesting(configPath, mockJSON, mockFileUtils)

	// Assert
	assert.Equal(t, expected, got)
	assert.Nil(t, err)
}

func TestGetConfigurationReadFailure(t *testing.T) {
	// Arrange
	mockJSON := new(mocks.DataUtils)
	mockFileUtils := new(mocks.FileUtils)

	expectedErr := errors.New("could not find file")
	expected := apifarm.Configuration{}

	mockFileUtils.On("Read", configPath).Return(nil, expectedErr)

	// Act
	got, gotErr := apifarm.GetConfigurationForTesting(configPath, mockJSON, mockFileUtils)

	// Assert
	assert.Equal(t, expected, got)
	assert.Equal(t, expectedErr, gotErr)
}

func TestGetConfigurationDeserializeFailure(t *testing.T) {
	// Arrange
	mockJSON := new(mocks.DataUtils)
	mockFileUtils := new(mocks.FileUtils)

	configData := []byte{24, 66, 80}
	expectedErr := errors.New("failed to deserialize")
	expected := apifarm.Configuration{}

	mockFileUtils.On("Read", configPath).Return(configData, nil)
	mockJSON.On("DeserializeConfiguration", configData).Return(&expected, expectedErr)

	// Act
	got, gotErr := apifarm.GetConfigurationForTesting(configPath, mockJSON, mockFileUtils)

	// Assert
	assert.Equal(t, expected, got)
	assert.Equal(t, expectedErr, gotErr)
}
