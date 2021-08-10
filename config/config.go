package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetConfig() RunMode {
	file, _ := os.Open("config.json")
	file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v\n", configuration)

	if os.Args[1] == "master" {
		return configuration.Master
	} else {
		return configuration.Slave
	}
}
