package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

type GatewayConfig struct {
	URL     string `yaml:"url"`
	Enabled bool   `yaml:"enabled"`
}

type Config struct {
	Gateways    map[string]GatewayConfig `yaml:"gateways"`
	Middlewares []string                 `yaml:"middlewares"`
	Static      struct {
		APIVersion            string `yaml:"apiVersion"`
		ServiceName           string `yaml:"serviceName"`
		DefaultTimeoutSeconds int    `yaml:"defaultTimeoutSeconds"`
		Host                  string `yaml:"host"`
		Port                  int    `yaml:"port"`
	} `yaml:"static"`
}

var (
	config     *Config
	configOnce sync.Once
)

func GetConfig() *Config {
	configOnce.Do(func() {
		path := filepath.Join("internal", "config", "config.yaml")
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("failed to read config file: %v", err)
		}
		cfg := &Config{}
		if err := yaml.Unmarshal(data, cfg); err != nil {
			log.Fatalf("failed to unmarshal config: %v", err)
		}
		// Dynamically update gateway URLs
		host := cfg.Static.Host
		port := cfg.Static.Port
		for name, gw := range cfg.Gateways {
			url := gw.URL
			url = strings.ReplaceAll(url, "{host}", host)
			url = strings.ReplaceAll(url, "{port}", fmt.Sprintf("%d", port))
			gw.URL = url
			cfg.Gateways[name] = gw
		}
		config = cfg
	})
	return config
}
