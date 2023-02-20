package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load .env file")
	}

	// Create a new Config struct with default values
	config := &Config{
		ServerPort:       "8080",
		DatabaseHost:     "localhost",
		DatabasePort:     "5432",
		DatabaseName:     "",
		DatabaseUser:     "",
		DatabasePassword: "",
	}

	// Set attributes for Config struct
	for _, attribute := range (&Config{}).getAttrName() {
		if val, ok := os.LookupEnv(attribute); ok {
			config = config.setProperty(attribute, val)
		}
	}

	Cfg = config
}
