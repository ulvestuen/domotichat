package config

import "os"

var (
	OpenAiApiKey string
	OpenAiOrg    string
)

func init() {
	OpenAiApiKey = os.Getenv("OPENAI_API_KEY")
	OpenAiOrg = os.Getenv("OPENAI_ORG")
}
