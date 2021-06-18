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
	b, err := loader.f.Read(p)

	if err != nil {
		return loader.qf.Error(err)
	}

	vgs, err := loader.json.DeserializeVideoGames(b)

	if err != nil {
		return loader.qf.Error(err)
	}

	for _, vg := range *vgs {
		loader.storage.AddVideoGame(vg)
	}

	return loader.qf.BuildMessage(SuccessfullyLoadedData, uint(0))
}
