package apifarm_test

import (
	mocks "apifarm/mocks/src"
	apifarm "apifarm/src"
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
