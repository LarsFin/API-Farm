package apifarm

import "encoding/json"

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
	var vg *VideoGame
	err := json.Unmarshal(data, vg)
	return vg, err
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
