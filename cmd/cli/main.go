package main

import (
	"github.com/Hossara/linkin-chat/cli/pkg/utils"
	"github.com/Hossara/linkin-chat/config"
	"github.com/spf13/cobra"
	"log"
	"path"

	"github.com/Hossara/linkin-chat/cli"
)

func main() {
	cobra.OnInitialize(initConfig)
	cli.Execute()
}

func initConfig() {
	home, err := utils.GetStoragePath("linkin-chat")

	if err != nil {
		log.Fatalf("Error getting home storage: %v", err)
	}

	err = config.ReadConfig(path.Join(home, "config.json"))

	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
}
