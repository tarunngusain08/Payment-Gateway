package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

type CircuitBreakerConfig struct {
	Enabled      bool    `yaml:"enabled"`
	MaxRequests  uint32  `yaml:"maxRequests"`
	Interval     int     `yaml:"intervalSeconds"`
	Timeout      int     `yaml:"timeoutSeconds"`
	FailureRatio float64 `yaml:"failureRatio"`
}

type ResilienceConfig struct {
	HTTPTimeoutSeconds   int                  `yaml:"httpTimeoutSeconds"`
	MaxRetries           int                  `yaml:"maxRetries"`
	InitialBackoffMillis int                  `yaml:"initialBackoffMillis"`
	MaxBackoffMillis     int                  `yaml:"maxBackoffMillis"`
	CircuitBreaker       CircuitBreakerConfig `yaml:"circuitBreaker"`
}

type GatewayConfig struct {
	URL     string `yaml:"url"`
	Enabled bool   `yaml:"enabled"`
	Name    string `yaml:"name,omitempty"` // Optional name for the gateway
}

type CacheConfig struct {
	InvalidationIntervalSeconds int `yaml:"invalidationIntervalSeconds"`
	TTLSeconds                  int `yaml:"ttlSeconds"`
}

type WorkerPoolConfig struct {
	NumWorkers int `yaml:"numWorkers"`
}

type Config struct {
	Gateways    map[string]GatewayConfig `yaml:"gateways"`
	Middlewares []string                 `yaml:"middlewares"`
	Static      struct {
		APIVersion            string `yaml:"apiVersion"`
		ServiceName           string `yaml:"serviceName"`
		DefaultTimeoutSeconds int    `yaml:"defaultTimeoutSeconds"`
		GatewayTimeoutSeconds int    `yaml:"gatewayTimeoutSeconds"`
		Host                  string `yaml:"host"`
		Port                  int    `yaml:"port"`
	} `yaml:"static"`
	Resilience ResilienceConfig `yaml:"resilience"`
	Cache      CacheConfig      `yaml:"cache"`
	WorkerPool WorkerPoolConfig `yaml:"workerPool"`
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
