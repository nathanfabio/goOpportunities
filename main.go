package main

import (
	"github.com/nathanfabio/goOpportunities/config"
	"github.com/nathanfabio/goOpportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}

	//Initialize  router
	router.Initialize()
}
