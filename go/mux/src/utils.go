package apifarm

import "encoding/json"

type DataUtils interface {
	Serialize(interface{}) ([]byte, error)
}

type Json struct {
}

func (*Json) Serialize(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
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
