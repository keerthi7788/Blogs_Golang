package config

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v2"
)

var (
	db *mongo.Database
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		URI  string `yaml:"uri"`
		Name string `yaml:"name"`
	} `yaml:"database"`
	JWT struct {
		Secret          string `yaml:"secret"`
		ExpirationHours int    `yaml:"expiration_hours"`
	} `yaml:"jwt"`
}

func LoadConfig() *Config {
	file, err := os.ReadFile("../../config/config.yml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}

	return &cfg
}

// GetDB returns the MongoDB database instance
// func GetDB() *mongo.Database {
// 	if db == nil {
// 		ConnectDB()
// 	}
// 	return db
// }

// func ConnectDB() {
// 	panic("unimplemented")
// }
