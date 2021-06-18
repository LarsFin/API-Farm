package apifarm

type DataLoader interface {
	Load(string) []VideoGame
}

type JSONFileLoader struct {
	json DataUtils
	f    FileUtils
}

func NewJSONFileLoader() *JSONFileLoader {
	return &JSONFileLoader{
		&JSON{},
		&fileUtils{},
	}
}

func NewJSONFileLoaderWithUtils(json DataUtils, f FileUtils) *JSONFileLoader {
	return &JSONFileLoader{
		json,
		f,
	}
}

func (loader *JSONFileLoader) Load(p string) []VideoGame {
	b, _ := loader.f.Read(p)

	vgs, _ := loader.json.DeserializeVideoGames(b)

	return *vgs
}
