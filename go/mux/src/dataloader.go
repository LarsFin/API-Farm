package apifarm

type DataLoader interface {
	Load(string) Query
}

type JSONFileLoader struct {
	storage DB
	json    DataUtils
	f       FileUtils
	qf      QueryFactory
}

func NewJSONFileLoader(storage DB) *JSONFileLoader {
	return &JSONFileLoader{
		storage,
		&JSON{},
		&fileUtils{},
		&queryFactory{},
	}
}

func NewJSONFileLoaderWithUtils(storage DB, json DataUtils, f FileUtils, qf QueryFactory) *JSONFileLoader {
	return &JSONFileLoader{
		storage,
		json,
		f,
		qf,
	}
}

func (loader *JSONFileLoader) Load(p string) Query {
	return Query{}
}
