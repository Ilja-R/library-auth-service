package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"

	"github.com/sethvargo/go-envconfig"

	"github.com/Ilja-R/library-auth-service/internal/bootstrap"
	"github.com/Ilja-R/library-auth-service/internal/config"
)

// @title AuthService API
// @contact.name AuthService API Service
// @contact.url http://test.com
// @contact.email test@test.com
func main() {

	// Load .env file into environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}

	var cfg config.Config

	err := envconfig.ProcessWith(context.TODO(), &envconfig.Config{Target: &cfg, Lookuper: envconfig.OsLookuper()})
	if err != nil {
		panic(err)
	}

	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, os.Interrupt)

	app := bootstrap.New(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-quitSignal
		cancel()
	}()

	app.Run(ctx)
}
