package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

// Config holds application configuration
type Config struct {
	// ADBPath is the path to the ADB executable
	ADBPath string
	
	// LogLevel is the logging level (debug, info, warn, error)
	LogLevel string
	
	// DefaultOutputFormat is the default JSON output format (pretty, compact)
	DefaultOutputFormat string
	
	// Timeout is the default timeout for ADB commands in seconds
	Timeout int
	
	// Platform is the detected platform (darwin, linux, windows)
	Platform string
}

var (
	// Global config instance
	globalConfig *Config
	once         sync.Once
)

// Load loads configuration from environment variables and defaults
func Load() *Config {
	once.Do(func() {
		globalConfig = &Config{
			ADBPath:              getEnv("ADBJSON_ADB_PATH", ""),
			LogLevel:             getEnv("ADBJSON_LOG_LEVEL", "info"),
			DefaultOutputFormat:  getEnv("ADBJSON_OUTPUT_FORMAT", "pretty"),
			Timeout:              getEnvInt("ADBJSON_TIMEOUT", 30),
			Platform:             detectPlatform(),
		}
	})
	return globalConfig
}

// Get returns the global config instance
func Get() *Config {
	if globalConfig == nil {
		return Load()
	}
	return globalConfig
}

// getEnv returns the environment variable or default value
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

// getEnvInt returns the environment variable as int or default value
func getEnvInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		var intVal int
		if _, err := fmt.Sscanf(val, "%d", &intVal); err == nil {
			return intVal
		}
	}
	return defaultVal
}

// detectPlatform detects the current platform
func detectPlatform() string {
	return runtime.GOOS
}

// GetConfigPath returns the path to the config file
func GetConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homeDir, ".adbjson", "config.yaml")
}

// LoadFromFile loads configuration from a file (placeholder for future YAML support)
func LoadFromFile(path string) (*Config, error) {
	// TODO: Implement YAML config file loading
	return Load(), nil
}
