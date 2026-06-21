package config

// This file defines constants for environment variable names used in the configuration.
const (
	// EnvLogLevel is the environment variable name for the logging level.
	EnvLogLevel = "LOG_LEVEL"

	// EnvHTTPAddress is the environment variable name for the HTTP server address.
	EnvHTTPAddress = "HTTP_ADDRESS"

	// EnvHTTPReadTimeout is the environment variable name for the HTTP server read timeout.
	EnvHTTPReadTimeout = "HTTP_READ_TIMEOUT"

	// EnvHTTPReadHeaderTimeout is the environment variable name for the HTTP server read header timeout.
	EnvHTTPReadHeaderTimeout = "HTTP_READ_HEADER_TIMEOUT"

	// EnvHTTPWriteTimeout is the environment variable name for the HTTP server write timeout.
	EnvHTTPWriteTimeout = "HTTP_WRITE_TIMEOUT"

	// EnvHTTPIdleTimeout is the environment variable name for the HTTP server idle timeout.
	EnvHTTPIdleTimeout = "HTTP_IDLE_TIMEOUT"

	// EnvHTTPShutdownTimeout is the environment variable name for the HTTP server shutdown timeout.
	EnvHTTPShutdownTimeout = "HTTP_SHUTDOWN_TIMEOUT"
)
