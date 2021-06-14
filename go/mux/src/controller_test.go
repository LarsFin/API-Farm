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
	query := apifarm.Query{
		result,
		0,
		nil,
	}

	mockService.On("GetAll").Return(query)
	mockResponse.On("OkJson", result)

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
	query := apifarm.Query{
		nil,
		500,
		err,
	}

	mockService.On("GetAll").Return(query)
	mockResponse.On("Error", err)

	// Act
	subject.HandleGetAll(mockResponse)

	// Assert
	mockService.AssertExpectations(t)
	mockResponse.AssertExpectations(t)
}
