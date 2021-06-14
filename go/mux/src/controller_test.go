package apifarm_test

import (
	mock "apifarm/mock"
	apifarm "apifarm/src"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestHandlePing(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock.NewMockResponse(ctrl)

	c := apifarm.Controller{}

	// Assert
	m.EXPECT().OkText("pong")

	// Act
	c.HandlePing(m)
}
