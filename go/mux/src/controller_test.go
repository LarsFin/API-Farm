package apifarm_test

import (
	mock "apifarm/mock"
	apifarm "apifarm/src"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestHandlePing(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock.NewMockResponse(ctrl)

	c := apifarm.Controller{}

	m.EXPECT().OkText("pong")

	c.HandlePing(m)
}
