package main

import (
	"github.com/Hossara/linkin-chat/cli/config"
	"github.com/Hossara/linkin-chat/cli/pkg/utils"
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
}
