package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"sideEcho/api"
	"sideEcho/config"
	_ "sideEcho/docs"
)

// @title         SideEcho API
// @version       0.1
// @description   This is a sample Exchange server.
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
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
