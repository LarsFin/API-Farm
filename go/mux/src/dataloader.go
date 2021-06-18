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
	b, _ := loader.f.Read(p)

	vgs, _ := loader.json.DeserializeVideoGames(b)

	for _, vg := range *vgs {
		loader.storage.AddVideoGame(vg)
	}

	return loader.qf.BuildMessage(SuccessfullyLoadedData, uint(0))
}
