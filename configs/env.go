package configs

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

const DEFAULT_HTTP_PORT = "3000"
const DEFAULT_GRAPHQL_PORT = "3001"

type EnvApp struct {
	HTTP_PORT    string
	GRAPHQL_PORT string
}

func LoadEnv() EnvApp {
	httpPort := os.Getenv("HTTP_PORT")
	graphqlPort := os.Getenv("GRAPHQL_PORT")

	if httpPort == "" {
		httpPort = DEFAULT_HTTP_PORT
	}
	if graphqlPort == "" {
		graphqlPort = DEFAULT_HTTP_PORT
	}

	return EnvApp{
		HTTP_PORT:    httpPort,
		GRAPHQL_PORT: graphqlPort,
	}
}
