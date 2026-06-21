// Package config provides configuration structures for the application.
// It loads configuration from environment variables and makes it available to the rest of the application.
package config

import (
	"os"
	"time"
)

// Config is the main configuration structure for the application.
type Config struct {
	// Log holds configuration for the logging system.
	Log Logger
	// HTTP holds configuration for the HTTP server.
	HTTP HTTP
}

// Logger holds configuration for the logging system.
type Logger struct {
	// Level is the logging level (e.g., "debug", "info", "warn", "error").
	Level string
}

// HTTP holds configuration for the HTTP server.
type HTTP struct {
	// Address is the address the HTTP server will listen on (e.g., ":8080", "localhost:8080").
	Address string
	// ReadTimeout is the maximum duration for reading the entire request, including the body.
	ReadTimeout time.Duration
	// ReadHeaderTimeout is the maximum duration for reading the request headers.
	ReadHeaderTimeout time.Duration
	// WriteTimeout is the maximum duration before timing out writes of the response.
	WriteTimeout time.Duration
	// IdleTimeout is the maximum duration a connection can be idle before timing out.
	IdleTimeout time.Duration
	// ShutdownTimeout is the maximum duration to wait for active requests to complete during shutdown.
	ShutdownTimeout time.Duration
}

// Load loads the configuration from environment variables and returns a Config struct.
// It returns an error if any required configuration is missing or invalid.
func Load() (Config, error) {
	cfg := Config{
		Log: Logger{
			Level: getEnv(EnvLogLevel, "info"),
		},
		HTTP: HTTP{
			Address:           getEnv(EnvHTTPAddress, ":8080"),
			ReadTimeout:       getEnvAsDuration(EnvHTTPReadTimeout, 30*time.Second),
			ReadHeaderTimeout: getEnvAsDuration(EnvHTTPReadHeaderTimeout, 10*time.Second),
			WriteTimeout:      getEnvAsDuration(EnvHTTPWriteTimeout, 30*time.Second),
			IdleTimeout:       getEnvAsDuration(EnvHTTPIdleTimeout, 120*time.Second),
			ShutdownTimeout:   getEnvAsDuration(EnvHTTPShutdownTimeout, 30*time.Second),
		},
	}
	return cfg, nil
}

// getEnv retrieves the value of the environment variable named by the key.
// If the variable is present in the environment, its value is returned.
// Otherwise, the defaultValue is returned.
func getEnv(envVar string, defaultValue string) string {
	val, exists := os.LookupEnv(envVar)
	if !exists || val == "" {
		return defaultValue
	}
	return val
}

// getEnvAsDuration retrieves the value of the environment variable named by envVar and parses it as a time.Duration.
// If the variable is not set or cannot be parsed, the defaultValue is returned.
// The environment variable should be in a format accepted by time.ParseDuration (e.g., "30s", "1m", "500ms").
func getEnvAsDuration(envVar string, defaultValue time.Duration) time.Duration {
	val, exists := os.LookupEnv(envVar)
	if !exists || val == "" {
		return defaultValue
	}
	d, err := time.ParseDuration(val)
	if err != nil {
		return defaultValue
	}
	return d
}
