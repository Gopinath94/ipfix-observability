# IPFix Observability

## Configuration

Configuration is loaded from Environment Variables. supported variables:

| Env variable | Description |
| -------- | -------- |
| LOG_LEVEL   | Logging level, e.g. "debug", "info", "warn", "error"   |
| HTTP_ADDRESS   | HTTP server address, e.g. ":8080", "localhost:8080"   |
| HTTP_READ_TIMEOUT  | HTTP server read timeout, e.g. "30s", "1m", "500ms"  |
| HTTP_READ_HEADER_TIMEOUT   | HTTP server read header timeout, e.g. "30s", "1m", "500ms"   |
| HTTP_WRITE_TIMEOUT  | HTTP server write timeout, e.g. "30s", "1m", "500ms"   |
| HTTP_IDLE_TIMEOUT   | HTTP server idle timeout, e.g. "30s", "1m", "500ms"  |
| HTTP_SHUTDOWN_TIMEOUT   | HTTP server shutdown timeout, e.g. "30s", "1m", "500ms"   |
