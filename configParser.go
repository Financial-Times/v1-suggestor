package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kr/pretty"
)

// ParseConfig opens the file at configFileName and unmarshals it into an AppConfig.
func ParseConfig(configFileName string) (*AppConfig, error) {
	file, err := ioutil.ReadFile(configFileName)
	if err != nil {
		errorLogger.Printf("Error reading configuration file [%v]: [%v]", configFileName, err.Error())
		return nil, err
	}

	var conf AppConfig
	err = json.Unmarshal(file, &conf)
	if err != nil {
		errorLogger.Printf("Error unmarshalling configuration file [%v]: [%v]", configFileName, err.Error())
		return nil, err
	}

	infoLogger.Printf("Using configuration: %# v", pretty.Formatter(conf))
	return &conf, nil
}
