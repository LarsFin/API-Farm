package apifarm

import "encoding/json"

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

type QueryFactory interface {
	Build([]byte, uint) Query
	Error(error) Query
}

type queryFactory struct {
}

func (*queryFactory) Build(result []byte, code uint) Query {
	return Query{
		result,
		code,
		nil,
	}
}

func (*queryFactory) Error(err error) Query {
	return Query{
		nil,
		500,
		err,
	}
}

type Query struct {
	Result []byte
	Code   uint
	Error  error
}
