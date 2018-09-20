package client

import (
	"encoding/json"
	"os"
)

var ConfigPath string

type Gateway struct {
	Url    string `json:"url"`
	Client Client `json:"client"`
}

type Client struct {
	Key  string `json:"key"`
	Pass string `json:"pass"`
}

// Setting config path for config and using in http client
func SetConfigPath(path string) string {
	ConfigPath = path

	return ConfigPath
}

func Config() Gateway {

	var gateway Gateway
	// Open config file
	jsonFile, err := os.Open(ConfigPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}
	jsonParser := json.NewDecoder(jsonFile)
	// Close config file
	defer jsonFile.Close()
	if err = jsonParser.Decode(&gateway); err != nil {
		panic(err)
	}

	return gateway
}
