package apifarm

type Configuration struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

func GetConfiguration(p string, json DataUtils, f FileUtils) Configuration {
	return Configuration{}
}
