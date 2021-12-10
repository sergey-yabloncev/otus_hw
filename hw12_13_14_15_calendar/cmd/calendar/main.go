package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sergey-yabloncev/otus_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/sergey-yabloncev/otus_hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/sergey-yabloncev/otus_hw/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/sergey-yabloncev/otus_hw/hw12_13_14_15_calendar/internal/storage/memory"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config, err := NewConfig(configFile)
	if err != nil {
		log.Fatalf("can't get config: %v", err)
	}

	logg, err := logger.New(config.Logger.Level, config.Logger.Path)
	if err != nil {
		log.Fatalf("can't start logger: %v", err)
	}

	application := app.New(*logg, memorystorage.New(logg))

	server := internalhttp.NewServer(
		*logg,
		config.Server.Port,
		*internalhttp.NewHandler(
			*logg,
			application,
		),
	)

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP,
	)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
