package env

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	ExaAIAPIKey  string
	Port         string
	AssetVersion string
	Dev          bool
}

func ParseEnv() Env {
	_ = godotenv.Load()

	return Env{
		ExaAIAPIKey:  os.Getenv("EXA_AI_API_KEY"),
		Port:         getEnvOr("PORT", "3000"),
		AssetVersion: fmt.Sprintf("%d", time.Now().Unix()),
		Dev:          os.Getenv("DEV") == "true",
	}
}

func getEnvOr(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
