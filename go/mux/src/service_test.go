package apifarm_test

import (
	mocks "apifarm/mocks/src"
	apifarm "apifarm/src"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// VideoGameService -> GetAll

func TestVideoGameServiceGetAllSuccessful(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	storedVideoGames := []apifarm.VideoGame{{Name: "Lady's Quest 3"}}
	serializedVideoGames := []byte{42, 42, 42}
	expectedQuery := apifarm.Query{}

	mockStorage.On("GetAllVideoGames").Return(storedVideoGames)
	mockJSON.On("Serialize", storedVideoGames).Return(serializedVideoGames, nil)
	mockQueryFactory.On("Build", serializedVideoGames, uint(0)).Return(expectedQuery)

	// Act
	actualQuery := subject.GetAll()

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

func TestVideoGameServiceGetAllSerializationFailure(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	storedVideoGames := []apifarm.VideoGame{{Name: "Lady's Quest 3"}}
	err := errors.New("failed to serialize to json")
	expectedQuery := apifarm.Query{}

	mockStorage.On("GetAllVideoGames").Return(storedVideoGames)
	mockJSON.On("Serialize", storedVideoGames).Return(nil, err)
	mockQueryFactory.On("Error", err).Return(expectedQuery)

	// Act
	actualQuery := subject.GetAll()

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

// VideoGameService -> Add

func TestVideoGameServiceAddSuccessful(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	reqData := []byte{20, 18, 24}
	videoGame := apifarm.VideoGame{
		Name: "Lady's Quest 3", DateReleased: time.Now(),
	}
	serializedVideoGame := []byte{23, 19, 18}
	expectedQuery := apifarm.Query{}

	mockStorage.On("AddVideoGame")
	mockJSON.On("DeserializeVideoGame", reqData).Return(&videoGame, nil)
	mockJSON.On("Serialize", videoGame).Return(serializedVideoGame)
	mockQueryFactory.On("Build", serializedVideoGame, uint(0)).Return(expectedQuery)

	// Act
	actualQuery := subject.Add(reqData)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}
