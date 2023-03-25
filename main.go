package main

import (
	"github.com/MogLuiz/Gopportunities-api/config"
	"github.com/MogLuiz/Gopportunities-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
