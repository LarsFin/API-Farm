package apifarm

type Configuration struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

func GetConfiguration(p string, json DataUtils, f FileUtils) (Configuration, error) {
	b, err := f.Read(p)

	if err != nil {
		return Configuration{}, err
	}

	c, _ := json.DeserializeConfiguration(b)

	return *c, nil
}
