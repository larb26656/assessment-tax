package main

import (
	"fmt"

	"github.com/larb26656/assessment-tax/config"
	"github.com/larb26656/assessment-tax/server"
)

func main() {
	// read configuration
	appConfig, err := config.NewAppConfigFromEnv()

	if err != nil {
		panic(fmt.Sprintf("Failed to load config %s", err))
	}

	// start server
	server.InitServer(appConfig)
}
