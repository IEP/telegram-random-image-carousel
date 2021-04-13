package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	BotAPIToken string `yaml:"bot_api_token"`
	WebhookURL  string `yaml:"webhook_url"`
	Port        string `yaml:"port"`
}

// LoadConfig for bot operation. See config.example.yaml
func LoadConfig(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return &Config{}, err
	}
	defer f.Close()

	content, _ := ioutil.ReadAll(f)
	var cfg Config
	_ = yaml.Unmarshal(content, &cfg)

	return &cfg, nil
}
