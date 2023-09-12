package main

import (
	"github.com/paula-michele-brisa/API-Gopportunities/config"
	"github.com/paula-michele-brisa/API-Gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Inicializar Configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization erro: %v", err)
		return
	}

	// Inicializar Router
	router.Initialize()
}
