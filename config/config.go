package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func GetConfig() RunMode {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	if os.Args[1] == "master" {
		return configuration.Master
	} else {
		return configuration.Slave
	}
}
