package apifarm_test

import (
	mocks "apifarm/mocks/src"
	apifarm "apifarm/src"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// VideoGameService -> Get

func TestVideoGameServiceGetSuccessful(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	const id = uint(5)
	storedVideoGame := apifarm.VideoGame{Name: "Autotron 95"}
	serializedVideoGame := []byte{12, 44, 44}
	expectedQuery := apifarm.Query{}

	mockStorage.On("GetVideoGame", id).Return(&storedVideoGame)
	mockJSON.On("Serialize", &storedVideoGame).Return(serializedVideoGame, nil)
	mockQueryFactory.On("BuildResult", serializedVideoGame, uint(0)).Return(expectedQuery)

	// Act
	actualQuery := subject.Get(id)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

func TestVideoGameServiceGetNotFound(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	const id = uint(99)
	expectedQuery := apifarm.Query{}

	mockStorage.On("GetVideoGame", id).Return(nil)
	mockQueryFactory.On("BuildMessage", apifarm.VideoGameNotFound(id), uint(404)).Return(expectedQuery)

	// Act
	actualQuery := subject.Get(id)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

func TestVideoGameServiceGetSerializationFail(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	const id = uint(5)
	storedVideoGame := apifarm.VideoGame{Name: "Autotron 95"}
	err := errors.New("failed to serialize to json")
	expectedQuery := apifarm.Query{}

	mockStorage.On("GetVideoGame", id).Return(&storedVideoGame)
	mockJSON.On("Serialize", &storedVideoGame).Return(nil, err)
	mockQueryFactory.On("Error", err).Return(expectedQuery)

	// Act
	actualQuery := subject.Get(id)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

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
	videoGameName := "Lady's Quest 3"
	videoGameDevelopers := []string{"A", "B"}
	videoGameProducers := []string{"A"}
	videoGameDateReleased := apifarm.CustomTime{time.Now()}
	videoGame := apifarm.VideoGame{
		Name:         videoGameName,
		Developers:   videoGameDevelopers,
		Producers:    videoGameProducers,
		DateReleased: videoGameDateReleased,
	}
	defaultedVideoGame := apifarm.VideoGame{
		uint(0),
		videoGameName,
		videoGameDevelopers,
		[]string{},
		[]string{},
		videoGameProducers,
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		videoGameDateReleased,
	}
	storedVideoGame := apifarm.VideoGame{
		ID:           1,
		Name:         videoGameName,
		DateReleased: videoGameDateReleased,
	}
	serializedVideoGame := []byte{23, 19, 18}
	expectedQuery := apifarm.Query{}

	mockJSON.On("DeserializeVideoGame", reqData).Return(&videoGame, nil)
	mockStorage.On("AddVideoGame", defaultedVideoGame).Return(storedVideoGame)
	mockJSON.On("Serialize", &storedVideoGame).Return(serializedVideoGame, nil)
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
	err := &time.ParseError{Value: invalidTime}
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

func TestVideoGameServiceAddInvalidAttributeFailure(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	attribute := "testers"
	reqData := []byte{87, 20, 56}
	err := &apifarm.InvalidAttributeError{Attribute: attribute}
	expectedQuery := apifarm.Query{}

	mockJSON.On("DeserializeVideoGame", reqData).Return(nil, err)
	mockQueryFactory.On("BuildMessage", apifarm.VideoGameInvalidAttribute(attribute), uint(400)).Return(expectedQuery)

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
		uint(0),
		"Lady's Quest 3",
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		apifarm.CustomTime{time.Now()},
	}
	storedVideoGame := apifarm.VideoGame{
		ID:           1,
		Name:         "Lady's Quest 3",
		DateReleased: apifarm.CustomTime{time.Now()},
	}
	err := errors.New("failed to serialize to json")
	expectedQuery := apifarm.Query{}

	mockJSON.On("DeserializeVideoGame", reqData).Return(&videoGame, nil)
	mockStorage.On("AddVideoGame", videoGame).Return(storedVideoGame)
	mockJSON.On("Serialize", &storedVideoGame).Return(nil, err)
	mockQueryFactory.On("Error", err).Return(expectedQuery)

	// Act
	actualQuery := subject.Add(reqData)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

func TestVideoGameServiceUpdateSuccessful(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	id := uint(5)
	reqData := []byte{13, 34, 22}
	videoGameToUpdateWith := apifarm.VideoGame{
		Name:       "Updated Name",
		Developers: []string{"A", "B"},
		Designers:  []string{"A", "B", "C"},
		Artists:    []string{"A", "B", "C"},
	}
	videoGameToUpdate := apifarm.VideoGame{
		uint(5),
		"Old Name",
		[]string{"1", "2"},
		[]string{"1"},
		[]string{"1", "2"},
		[]string{"1", "2", "3"},
		[]string{"1", "2"},
		[]string{"1", "2", "3"},
		[]string{"1", "2"},
		[]string{"1", "2"},
		[]string{"1", "2", "3"},
		apifarm.CustomTime{Time: time.Now()},
	}
	updatedVideoGame := apifarm.VideoGame{
		videoGameToUpdate.ID,
		videoGameToUpdateWith.Name,
		videoGameToUpdateWith.Developers,
		videoGameToUpdate.Publishers,
		videoGameToUpdate.Directors,
		videoGameToUpdate.Producers,
		videoGameToUpdateWith.Designers,
		videoGameToUpdate.Programmers,
		videoGameToUpdateWith.Artists,
		videoGameToUpdate.Composers,
		videoGameToUpdate.Platforms,
		videoGameToUpdate.DateReleased,
	}
	serializedUpdatedVideoGame := []byte{12, 21, 35}
	expectedQuery := apifarm.Query{}

	mockStorage.On("GetVideoGame", id).Return(&videoGameToUpdate)
	mockJSON.On("DeserializeVideoGame", reqData).Return(&videoGameToUpdateWith, nil)
	mockStorage.On("UpdateVideoGame", updatedVideoGame).Return(&updatedVideoGame)
	mockJSON.On("Serialize", &updatedVideoGame).Return(serializedUpdatedVideoGame, nil)
	mockQueryFactory.On("BuildResult", serializedUpdatedVideoGame, uint(0)).Return(expectedQuery)

	// Act
	actualQuery := subject.Update(id, reqData)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

func TestVideoGameServiceUpdateNotFound(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	id := uint(5)
	expectedQuery := apifarm.Query{}

	mockStorage.On("GetVideoGame", id).Return(nil)
	mockQueryFactory.On("BuildMessage", apifarm.VideoGameNotFound(id), uint(404)).Return(expectedQuery)

	// Act
	actualQuery := subject.Update(id, []byte{})

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}

func TestVideoGameServiceUpdateInvalidDateFailure(t *testing.T) {
	// Arrange
	mockStorage := new(mocks.DB)
	mockJSON := new(mocks.DataUtils)
	mockQueryFactory := new(mocks.QueryFactory)

	subject := apifarm.NewVideoGameServiceWithUtils(mockStorage, mockJSON, mockQueryFactory)

	id := uint(5)
	reqData := []byte{13, 34, 22}
	invalidDate := "2010/08/26"
	videoGameToUpdate := apifarm.VideoGame{}
	err := time.ParseError{Value: invalidDate}
	expectedQuery := apifarm.Query{}

	mockStorage.On("GetVideoGame", id).Return(&videoGameToUpdate)
	mockJSON.On("DeserializeVideoGame", reqData).Return(nil, err)
	mockQueryFactory.On("BuildMessage", apifarm.VideoGameInvalidDate(invalidDate), uint(400)).Return(expectedQuery)

	// Act
	actualQuery := subject.Update(id, reqData)

	// Assert
	assert.Equal(t, expectedQuery, actualQuery)
	mockStorage.AssertExpectations(t)
	mockJSON.AssertExpectations(t)
	mockQueryFactory.AssertExpectations(t)
}
