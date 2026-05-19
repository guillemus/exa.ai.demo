package main

import (
	"exa.ai.demo/env"
	"exa.ai.demo/logger"
	"exa.ai.demo/server"
)

func main() {
	env := env.ParseEnv()
	logger.Setup(env)
	server.StartServer(env)
}
