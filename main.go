package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"sideEcho/api"
	"sideEcho/config"
)

func main() {
	flags := config.ReadFlags()
	cfg, err := config.ReadConfigFile(flags)
	if err != nil {
		panic(err)
	}
	cleanup := api.NewAPI(cfg)

	// sigint 받으면 shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Config.Api.ShutDownTimeoutSec)*time.Second)
	defer cancel()
	cleanup(ctx)
}
