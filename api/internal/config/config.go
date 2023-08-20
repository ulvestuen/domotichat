package config

import "os"

var (
	OpenAiApiKey string
	OpenAiOrg    string
	Port         string
)

func init() {
	OpenAiApiKey = os.Getenv("OPENAI_API_KEY")
	OpenAiOrg = os.Getenv("OPENAI_ORG")
	Port = GetOrDefault("PORT", "8080")
}

func GetOrDefault(key string, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}
