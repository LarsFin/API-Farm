package apifarm_test

import (
	mocks "apifarm/mocks/src"
	apifarm "apifarm/src"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// VideoGameService -> GetAll

func TestVideoGameServiceGetAllSuccessful(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJson := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJson, mockQueryFactory)

	storedVideoGames := []apifarm.VideoGame{{Name: "Lady's Quest 3"}}
	serializedVideoGames := []byte{42, 42, 42}
	expectedQuery := apifarm.Query{}

	mockStorage.On("GetAllVideoGames").Return(storedVideoGames)
	mockJson.On("Serialize", storedVideoGames).Return(serializedVideoGames, nil)
	mockQueryFactory.On("Build", serializedVideoGames, uint(0)).Return(expectedQuery)

	// Act
	actualQuery := subject.GetAll()

	// Assert
	assert.Equal(t, expectedQuery, actualQuery, "they should be equal")
	mockStorage.AssertExpectations(t)
	mockJson.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

func TestVideoGameServiceGetAllSerializationFailure(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJson := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJson, mockQueryFactory)

	storedVideoGames := []apifarm.VideoGame{{Name: "Lady's Quest 3"}}
	err := errors.New("Failed to serialize Json!")
	expectedQuery := apifarm.Query{}

	mockStorage.On("GetAllVideoGames").Return(storedVideoGames)
	mockJson.On("Serialize", storedVideoGames).Return(nil, err)
	mockQueryFactory.On("Error", err).Return(expectedQuery)

	// Act
	actualQuery := subject.GetAll()

	// Assert
	assert.Equal(t, expectedQuery, actualQuery, "they should be equal")
	mockStorage.AssertExpectations(t)
	mockJson.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}
