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
	mockRequest := new(mocks.Request)
	mockResponse := new(mocks.Response)

	subject := apifarm.NewController(nil)

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
