package main

// loads and parses the JSON config file

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	URLs []string `json:"urls"`
}

// LoadConfig reads a JSON configuration file and unmarshals it into a Config struct.
// It returns a pointer to the Config struct and an error if any occurs during file opening or JSON decoding.
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("could not decode config file: %w", err)
	}
	return config, nil
}
