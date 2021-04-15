package main

import (
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	reference := &Config{
		BotAPIToken: "1234567890",
		WebhookURL:  "https://some_website.com",
		Port:        "3000",
	}

	cfg, err := LoadConfig("config.example.yaml")
	if err != nil {
		t.Fatalf("config cannot be loaded %s", err)
	}

	if !reflect.DeepEqual(reference, cfg) {
		t.Fatalf("reference != config; %+v", cfg)
	}
}

func TestNonExistingConfig(t *testing.T) {
	cfg, err := LoadConfig("not_exist.yaml")
	if err == nil {
		t.Fatalf("LoadConfig supposed to return error")
	}
	if !reflect.DeepEqual(cfg, &Config{}) {
		t.Fatalf("loaded config must equal to default Config")
	}
}
