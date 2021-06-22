package apifarm_test

import (
	mocks "apifarm/mocks/src"
	apifarm "apifarm/src"
	"errors"
	"testing"
)

func TestHandlePing(t *testing.T) {
	// Arrange
	mockResponse := new(mocks.Response)
	mockResponse.On("OkText", "pong")

	subject := apifarm.Controller{}

	// Act
	subject.HandlePing(mockResponse)

	// Assert
	mockResponse.AssertExpectations(t)
}

func TestHandleGet200(t *testing.T) {
	// Arrange
	mockService := new(mocks.Service)
	mockRequest := new(mocks.Request)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(mockService)

	result := []byte{5, 16, 49}
	query := apifarm.Query{Result: result}

	mockRequest.On("GetParam", "id").Return("5")
	mockService.On("Get", uint(5)).Return(query)
	mockResponse.On("OkJSON", result)

	// Act
	subject.HandleGet(mockRequest, mockResponse)

	// Assert
	mockService.AssertExpectations(t)
	mockRequest.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandleGet400(t *testing.T) {
	// Arrange
	mockService := new(mocks.Service)
	mockRequest := new(mocks.Request)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(mockService)

	invalidID := "invalid!"

	mockRequest.On("GetParam", "id").Return(invalidID)
	mockResponse.On("BadRequestText", apifarm.ParamInvalidID(invalidID))

	// Act
	subject.HandleGet(mockRequest, mockResponse)

	// Assert
	mockRequest.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandleGet404(t *testing.T) {
	// Arrange
	mockService := new(mocks.Service)
	mockRequest := new(mocks.Request)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(mockService)

	queryMessage := "VIDEO GAME NOT FOUND"
	query := apifarm.Query{Message: queryMessage, Code: uint(404)}

	mockRequest.On("GetParam", "id").Return("99")
	mockService.On("Get", uint(99)).Return(query)
	mockResponse.On("NotFoundText", queryMessage)

	// Act
	subject.HandleGet(mockRequest, mockResponse)

	// Assert
	mockService.AssertExpectations(t)
	mockRequest.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandleGet500(t *testing.T) {
	// Arrange
	mockService := new(mocks.Service)
	mockRequest := new(mocks.Request)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(mockService)

	err := errors.New("query failed")
	query := apifarm.Query{Code: uint(500), Error: err}

	mockRequest.On("GetParam", "id").Return("5")
	mockService.On("Get", uint(5)).Return(query)
	mockResponse.On("Error", err)

	// Act
	subject.HandleGet(mockRequest, mockResponse)

	// Assert
	mockService.AssertExpectations(t)
	mockRequest.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandleGetAll200(t *testing.T) {
	// Arrange
	mockService := new(mocks.Service)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(mockService)

	result := []byte{12, 36, 18}
	query := apifarm.Query{Result: result}

	mockService.On("GetAll").Return(query)
	mockResponse.On("OkJSON", result)

	// Act
	subject.HandleGetAll(mockResponse)

	// Assert
	mockService.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandleGetAll500(t *testing.T) {
	// Arrange
	mockService := new(mocks.Service)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(mockService)

	err := errors.New("query failed")
	query := apifarm.Query{Code: uint(500), Error: err}

	mockService.On("GetAll").Return(query)
	mockResponse.On("Error", err)

	// Act
	subject.HandleGetAll(mockResponse)

	// Assert
	mockService.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandlePost201(t *testing.T) {
	// Arrange
	mockService := new(mocks.Service)
	mockRequest := new(mocks.Request)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(mockService)

	body := []byte{33, 12, 48}
	result := []byte{40, 23, 98}
	query := apifarm.Query{Result: result}

	mockRequest.On("GetBody").Return(body, nil)
	mockService.On("Add", body).Return(query)
	mockResponse.On("CreatedJSON", result)

	// Act
	subject.HandlePost(mockRequest, mockResponse)

	// Assert
	mockService.AssertExpectations(t)
	mockRequest.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandlePost400(t *testing.T) {
	// Arrange
	mockService := new(mocks.Service)
	mockRequest := new(mocks.Request)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(mockService)

	body := []byte{33, 12, 48}
	queryMessage := "INVALID VIDEO GAME"
	query := apifarm.Query{Message: queryMessage, Code: uint(400)}

	mockRequest.On("GetBody").Return(body, nil)
	mockService.On("Add", body).Return(query)
	mockResponse.On("BadRequestText", queryMessage)

	// Act
	subject.HandlePost(mockRequest, mockResponse)

	// Assert
	mockService.AssertExpectations(t)
	mockRequest.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandlePost500BodyReadFailure(t *testing.T) {
	// Arrange
	mockService := new(mocks.Service)
	mockRequest := new(mocks.Request)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(mockService)

	err := errors.New("body read failed")

	mockRequest.On("GetBody").Return(nil, err)
	mockResponse.On("Error", err)

	// Act
	subject.HandlePost(mockRequest, mockResponse)

	// Assert
	mockRequest.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandlePost500ServiceFailure(t *testing.T) {
	// Arrange
	mockService := new(mocks.Service)
	mockRequest := new(mocks.Request)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(mockService)

	body := []byte{33, 12, 48}
	err := errors.New("query failed")
	query := apifarm.Query{Code: uint(500), Error: err}

	mockRequest.On("GetBody").Return(body, nil)
	mockService.On("Add", body).Return(query)
	mockResponse.On("Error", err)

	// Act
	subject.HandlePost(mockRequest, mockResponse)

	// Assert
	mockService.AssertExpectations(t)
	mockRequest.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandlePut200(t *testing.T) {
	// Arrange
	mockService := new(mocks.Service)
	mockRequest := new(mocks.Request)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(mockService)

	body := []byte{20, 18, 48}
	result := []byte{90, 92, 56}
	query := apifarm.Query{Result: result}

	mockRequest.On("GetParam", "id").Return("5")
	mockRequest.On("GetBody").Return(body, nil)
	mockService.On("Update", uint(5), body).Return(query)
	mockResponse.On("OkJSON", result)

	// Act
	subject.HandlePut(mockRequest, mockResponse)

	// Assert
	mockService.AssertExpectations(t)
	mockRequest.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandlePut400InvalidID(t *testing.T) {
	// Arrange
	mockService := new(mocks.Service)
	mockRequest := new(mocks.Request)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(mockService)

	invalidID := "invalid!"

	mockRequest.On("GetParam", "id").Return(invalidID)
	mockResponse.On("BadRequestText", apifarm.ParamInvalidID(invalidID))

	// Act
	subject.HandlePut(mockRequest, mockResponse)

	// Assert
	mockService.AssertExpectations(t)
	mockRequest.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandlePut400FromQuery(t *testing.T) {
	// Arrange
	mockService := new(mocks.Service)
	mockRequest := new(mocks.Request)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(mockService)

	body := []byte{20, 18, 48}
	queryMessage := "INVALID VIDEO GAME"
	query := apifarm.Query{Message: queryMessage, Code: uint(400)}

	mockRequest.On("GetParam", "id").Return("5")
	mockRequest.On("GetBody").Return(body, nil)
	mockService.On("Update", uint(5), body).Return(query)
	mockResponse.On("BadRequestText", queryMessage)

	// Act
	subject.HandlePut(mockRequest, mockResponse)

	// Assert
	mockService.AssertExpectations(t)
	mockRequest.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandleTestSetup200(t *testing.T) {
	// Arrange
	mockDataLoader := new(mocks.DataLoader)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewAPITestingController(mockDataLoader)

	message := "SUCCESS MESSAGE"
	query := apifarm.Query{Message: message}

	mockDataLoader.On("Load", apifarm.SampleDataPath).Return(query)
	mockResponse.On("OkText", message)

	// Act
	subject.HandleTestSetup(mockResponse)

	// Assert
	mockDataLoader.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}

func TestHandleTestSetup500(t *testing.T) {
	// Arrange
	mockDataLoader := new(mocks.DataLoader)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewAPITestingController(mockDataLoader)

	err := errors.New("query failed")
	query := apifarm.Query{Code: uint(500), Error: err}

	mockDataLoader.On("Load", apifarm.SampleDataPath).Return(query)
	mockResponse.On("Error", err)

	// Act
	subject.HandleTestSetup(mockResponse)

	// Assert
	mockDataLoader.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}
