package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/larb26656/assessment-tax/config"
	"github.com/larb26656/assessment-tax/database"
	"github.com/larb26656/assessment-tax/server"
)

func main() {
	// read configuration
	appConfig, err := config.NewAppConfigFromEnv()

	if err != nil {
		panic(fmt.Sprintf("Failed to load config : %s", err))
	}

	// init database
	db, err := database.InitDatabase(appConfig)

	if err != nil {
		panic(fmt.Sprintf("Failed to initialize database : %s", err))
	}

	// init server
	e := server.InitServer(appConfig, db)

	go func() {
		if err := e.Start(":2565"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	log.Println("Close application")
}
