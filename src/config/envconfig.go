package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

type (
	// Configuration is struct for holding service's configuration info
	Configuration struct {
		MongoURL       string       `json:"MongoURL" validate:"required"`
		DatabaseName   string       `json:"DatabaseName" validate:"required"`
		ListenPort     string       `json:"ListenPort" validate:"required"`
		SeedsFilePaths Seeds        `json:"SeedsFilePaths" validate:"required"`
		Log            LoggerConfig `json:"Log" validate:"required"`
		UserGamesCount int          `json:"UserGamesCount", validate:"required"`
		WithSeeds      bool         `json:"WithSeeds", validate:"required"`
	}

	// LoggerConfig is a struct for holding logger configuration
	LoggerConfig struct {
		Level       uint32 `json:"Level" validate:"required"`
		ServiceName string `json:"ServiceName" validate:"required"`
		FileName    string `json:"FileName" validate:"required"`
	}

	Seeds struct {
		Users     string `json:"Users" validate:"required"`
		UserGames string `json:"UserGames" validate:"required"`
	}
)

func Load(configFilePath string) (err error, config Configuration) {
	if err, config = readConfigJSON(configFilePath); err != nil {
		return
	}

	return
}

// readConfigJSON reads config info from JSON file
func readConfigJSON(filePath string) (error, Configuration) {
	log.Printf("Searching JSON config file (%s)", filePath)
	var config Configuration

	contents, err := ioutil.ReadFile(filePath)
	if err == nil {
		reader := bytes.NewBuffer(contents)
		err = json.NewDecoder(reader).Decode(&config)
	}
	if err != nil {
		log.Printf("Error while reading configuration from JSON (%s) error: %s\n", filePath, err.Error())
	} else {
		log.Printf("Configuration from JSON (%s) provided\n", filePath)
	}

	return err, config
}
