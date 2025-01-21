package main

import (
	"flag"
	"github.com/Hossara/linkin-chat/app"
	"github.com/Hossara/linkin-chat/config"
	"github.com/Hossara/linkin-chat/server/http"
	"os"
)

var configPath = flag.String("config", "config.json", "Path to service config file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	appContainer := app.MustNewApp(c)

	err := http.Bootstrap(appContainer, c.Server)

	if err != nil {
		return
	}

}
