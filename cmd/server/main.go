package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Hossara/linkin-chat/app"
	"github.com/Hossara/linkin-chat/config"
	"github.com/Hossara/linkin-chat/server/http"
	"github.com/Hossara/linkin-chat/server/nats"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var configPath = flag.String("config", "config.json", "Path to service config file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	appContainer := app.MustNewApp(c)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	wg := sync.WaitGroup{}
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := http.Bootstrap(appContainer, c.Server)

		if err != nil {
			println(err.Error())
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := nats.Bootstrap(appContainer, c.Nats)

		if err != nil {
			panic(err.Error())
			return
		}
	}()

	<-ctx.Done()
	fmt.Println("Server received shutdown signal, waiting for components to stop...")
	return
}
