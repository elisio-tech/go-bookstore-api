package config

import (
	"os"
	"strings"
)

type Config struct {
	Port           string
	JWTSecret      string
	DBPath         string
	AllowedOrigins []string
	RateLimitReq   int
	RateLimitWin   int
}

var cfg *Config

func Load() *Config {
	cfg = &Config{
		Port:           getEnv("PORT", "3000"),
		JWTSecret:      getEnv("JWT_SECRET", "dev-insecure-secret-change-me"),
		DBPath:         getEnv("DB_PATH", "data/bookstore.db"),
		AllowedOrigins: splitEnv("ALLOWED_ORIGINS", "*"),
		RateLimitReq:   30,
		RateLimitWin:   60,
	}
	return cfg
}

func Get() *Config {
	if cfg == nil {
		return Load()
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func splitEnv(key, fallback string) []string {
	raw := getEnv(key, fallback)
	if raw == "" {
		return []string{}
	}
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}