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
	mockQueryFactory.On("BuildResult", serializedVideoGames, uint(0)).Return(expectedQuery)

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
		Name:         "Lady's Quest 3",
		DateReleased: apifarm.CustomTime{time.Now()},
	}
	storedVideoGame := apifarm.VideoGame{
		Id:           1,
		Name:         "Lady's Quest 3",
		DateReleased: apifarm.CustomTime{time.Now()},
	}
	serializedVideoGame := []byte{23, 19, 18}
	expectedQuery := apifarm.Query{}

	mockJSON.On("DeserializeVideoGame", reqData).Return(&videoGame, nil)
	mockStorage.On("AddVideoGame", videoGame).Return(storedVideoGame)
	mockJSON.On("Serialize", storedVideoGame).Return(serializedVideoGame, nil)
	mockQueryFactory.On("BuildResult", serializedVideoGame, uint(0)).Return(expectedQuery)

	// Act
	actualQuery := subject.Add(reqData)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

func TestVideoGameServiceAddInvalidDateFailure(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	invalidTime := "2006/01/02"
	reqData := []byte{28, 44, 21}
	err := &(time.ParseError{Value: invalidTime})
	expectedQuery := apifarm.Query{}

	mockJSON.On("DeserializeVideoGame", reqData).Return(nil, err)
	mockQueryFactory.On("BuildMessage", apifarm.VideoGameInvalidDate(invalidTime), uint(400)).Return(expectedQuery)

	// Act
	actualQuery := subject.Add(reqData)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

func TestVideoGameServiceAddDeserializationFailure(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	reqData := []byte{35, 34, 12}
	err := errors.New("failed to deserialize")
	expectedQuery := apifarm.Query{}

	mockJSON.On("DeserializeVideoGame", reqData).Return(nil, err)
	mockQueryFactory.On("BuildMessage", apifarm.InvalidJSON, uint(400)).Return(expectedQuery)

	// Act
	actualQuery := subject.Add(reqData)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

func TestVideoGameServiceAddNoNameFailure(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	reqData := []byte{20, 18, 24}
	videoGame := apifarm.VideoGame{
		DateReleased: apifarm.CustomTime{time.Now()},
	}
	expectedQuery := apifarm.Query{}

	mockJSON.On("DeserializeVideoGame", reqData).Return(&videoGame, nil)
	mockQueryFactory.On("BuildMessage", apifarm.VideoGameNameRequired, uint(400)).Return(expectedQuery)

	// Act
	actualQuery := subject.Add(reqData)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

func TestVideoGameServiceAddNoDateFailure(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	reqData := []byte{20, 18, 24}
	videoGame := apifarm.VideoGame{
		Name: "Lady's Quest 3",
	}
	expectedQuery := apifarm.Query{}

	mockJSON.On("DeserializeVideoGame", reqData).Return(&videoGame, nil)
	mockQueryFactory.On("BuildMessage", apifarm.VideoGameDateRequired, uint(400)).Return(expectedQuery)

	// Act
	actualQuery := subject.Add(reqData)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

func TestVideoGameServiceAddSerializationFailure(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	reqData := []byte{20, 18, 24}
	videoGame := apifarm.VideoGame{
		Name:         "Lady's Quest 3",
		DateReleased: apifarm.CustomTime{time.Now()},
	}
	storedVideoGame := apifarm.VideoGame{
		Id:           1,
		Name:         "Lady's Quest 3",
		DateReleased: apifarm.CustomTime{time.Now()},
	}
	err := errors.New("failed to serialize to json")
	expectedQuery := apifarm.Query{}

	mockJSON.On("DeserializeVideoGame", reqData).Return(&videoGame, nil)
	mockStorage.On("AddVideoGame", videoGame).Return(storedVideoGame)
	mockJSON.On("Serialize", storedVideoGame).Return(nil, err)
	mockQueryFactory.On("Error", err).Return(expectedQuery)

	// Act
	actualQuery := subject.Add(reqData)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}
