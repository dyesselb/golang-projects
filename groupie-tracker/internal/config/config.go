package config

import (
	"encoding/json"
	"os"
)


type Config struct{
	IsDebug bool `json:"is_debug"`
	Listen  struct {
		Protocol string `json:"protocol"`
		BindIp   string `json:"bind_ip"`
		Port     string `json:"port"`
	} `json:"listen"`
}

func LoadConfig(file string)(Config, error){
	var conf Config
	openedFile, err := os.Open(file)
	
	defer openedFile.Close()

	if err != nil{
		return conf, err
	}

	decodeNew := json.NewDecoder(openedFile)
	err = decodeNew.Decode(&conf)
	if err != nil{
		return conf, err
	}
	return conf, nil
}