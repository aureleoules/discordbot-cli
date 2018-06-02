package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//Pretty : Print JSON
func Pretty(data interface{}) {
	prettyfied, _ := json.MarshalIndent(data, "", "  ")
	log.Println(string(prettyfied))
}

//LoadConfig : Parse config file
func LoadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		return Config{}, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}
