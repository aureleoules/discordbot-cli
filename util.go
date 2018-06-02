package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func pretty(data interface{}, _ error) {
	prettyfied, _ := json.MarshalIndent(data, "", "\t")
	log.Println(string(prettyfied))
}

//LoadConfig : Parse config file
func LoadConfig(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
