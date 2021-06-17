package apifarm

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Custom Time

type CustomTime struct {
	time.Time
}

const layout = "02/01/2006"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	if s == "null" {
		ct.Time = time.Time{}
		return nil
	}

	t, err := time.Parse(layout, s)
	ct.Time = t

	return err
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("\"%s\"", ct.Time.Format(layout))

	return []byte(s), nil
}

// Data Utilities

type DataUtils interface {
	Serialize(interface{}) ([]byte, error)
	DeserializeVideoGame([]byte) (*VideoGame, error)
}

type JSON struct {
}

func (*JSON) Serialize(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func (*JSON) DeserializeVideoGame(data []byte) (*VideoGame, error) {
	var vg VideoGame
	err := json.Unmarshal(data, &vg)

	return &vg, err
}

// Factories

type QueryFactory interface {
	BuildResult([]byte, uint) Query
	BuildMessage(string, uint) Query
	Error(error) Query
}

type queryFactory struct {
}

func (*queryFactory) BuildResult(result []byte, code uint) Query {
	return Query{
		result,
		"Successfully built.",
		code,
		nil,
	}
}

func (*queryFactory) BuildMessage(msg string, code uint) Query {
	return Query{
		nil,
		msg,
		code,
		nil,
	}
}

func (*queryFactory) Error(err error) Query {
	return Query{
		nil,
		"An unforseen error occurred.",
		500,
		err,
	}
}

type Query struct {
	Result  []byte
	Message string
	Code    uint
	Error   error
}

// Messages

const VideoGameDateRequired = "a date_released is required for a video game."
const VideoGameNameRequired = "A name is required for a video game."
