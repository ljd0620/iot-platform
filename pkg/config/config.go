package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	MQTTBroker   string `json:"mqtt_broker"`
	WebSocketURL string `json:"websocket_url"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) GetConfig() *Config {
	return c
}