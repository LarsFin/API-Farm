package apifarm

type Configuration struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

func GetConfiguration(p string) (Configuration, error) {
	return getConfiguration(p, &JSON{}, &fileUtils{})
}

func GetConfigurationForTesting(p string, json DataUtils, f FileUtils) (Configuration, error) {
	return getConfiguration(p, json, f)
}

func getConfiguration(p string, json DataUtils, f FileUtils) (Configuration, error) {
	if json == nil {
		json = &JSON{}
	}

	if f == nil {
		f = &fileUtils{}
	}

	b, err := f.Read(p)

	if err != nil {
		return Configuration{}, err
	}

	c, err := json.DeserializeConfiguration(b)

	if err != nil {
		return Configuration{}, err
	}

	return *c, nil
}
