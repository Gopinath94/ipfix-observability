package config

import (
	"os"
	"testing"
	"time"
)

func TestLoadConfig(t *testing.T) {
	cases := []struct {
		desc string
		env  map[string]string
		want Config
	}{
		{
			desc: "Default configuration",
			env: map[string]string{
				EnvLogLevel:              "",
				EnvHTTPAddress:           "",
				EnvHTTPReadTimeout:       "",
				EnvHTTPReadHeaderTimeout: "",
				EnvHTTPWriteTimeout:      "",
				EnvHTTPIdleTimeout:       "",
				EnvHTTPShutdownTimeout:   "",
			},
			want: Config{
				Log: Logger{
					Level: "info",
				},
				HTTP: HTTP{
					Address:           ":8080",
					ReadTimeout:       30 * time.Second,
					ReadHeaderTimeout: 10 * time.Second,
					WriteTimeout:      30 * time.Second,
					IdleTimeout:       120 * time.Second,
					ShutdownTimeout:   30 * time.Second,
				},
			},
		},
		{
			desc: "Custom configuration from environment variables",
			env: map[string]string{
				EnvLogLevel:              "debug",
				EnvHTTPAddress:           "localhost:9090",
				EnvHTTPReadTimeout:       "15s",
				EnvHTTPReadHeaderTimeout: "5s",
				EnvHTTPWriteTimeout:      "20s",
				EnvHTTPIdleTimeout:       "60s",
				EnvHTTPShutdownTimeout:   "10s",
			},
			want: Config{
				Log: Logger{
					Level: "debug",
				},
				HTTP: HTTP{
					Address:           "localhost:9090",
					ReadTimeout:       15 * time.Second,
					ReadHeaderTimeout: 5 * time.Second,
					WriteTimeout:      20 * time.Second,
					IdleTimeout:       60 * time.Second,
					ShutdownTimeout:   10 * time.Second,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			for k, v := range c.env {
				t.Setenv(k, v)
			}

			got, err := Load()
			if err != nil {
				t.Fatalf("Load() returned an error: %v", err)
			}

			if got != c.want {
				t.Errorf("Load() = %+v; want %+v", got, c.want)
			}
		})
	}
}

func TestGetEnv(t *testing.T) {
	cases := []struct {
		desc string
		env  string
		val  string
		def  string
		want string
	}{
		{
			desc: "Environment variable is set",
			env:  "TEST_ENV_VAR",
			val:  "test_value",
			def:  "default_value",
			want: "test_value",
		},
		{
			desc: "Environment variable is not set, default value is used",
			env:  "TEST_ENV_VAR_NOT_SET",
			val:  "",
			def:  "default_value",
			want: "default_value",
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			if c.val != "" {
				t.Setenv(c.env, c.val)
			} else {
				os.Unsetenv(c.env)
			}

			got := getEnv(c.env, c.def)
			if got != c.want {
				t.Errorf("getEnv(%q, %q) = %q; want %q", c.env, c.def, got, c.want)
			}
		})
	}
}

func TestGetEnvAsDuration(t *testing.T) {
	cases := []struct {
		desc string
		env  string
		val  string
		def  time.Duration
		want time.Duration
	}{
		{
			desc: "Environment variable is set",
			env:  "TEST_ENV_VAR",
			val:  "50ms",
			def:  50 * time.Millisecond,
			want: 50 * time.Millisecond,
		},
		{
			desc: "Environment variable is not set, default value is used",
			env:  "TEST_ENV_VAR_NOT_SET",
			val:  "",
			def:  10 * time.Millisecond,
			want: 10 * time.Millisecond,
		},
		{
			desc: "Environment variable is set to an invalid duration, default value is used",
			env:  "TEST_ENV_VAR_INVALID",
			val:  "invalid_duration",
			def:  20 * time.Millisecond,
			want: 20 * time.Millisecond,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			if c.val != "" {
				t.Setenv(c.env, c.val)
			} else {
				os.Unsetenv(c.env)
			}

			got := getEnvAsDuration(c.env, c.def)
			if got != c.want {
				t.Errorf("getEnvAsDuration(%q, %q) = %q; want %q", c.env, c.def, got, c.want)
			}
		})
	}
}
