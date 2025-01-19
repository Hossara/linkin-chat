package main

import (
	"github.com/spf13/cobra"

	"github.com/Hossara/linkin-chat/cli"
)

func main() {
	cobra.OnInitialize(initConfig)
	cli.Execute()
}

func initConfig() {

}
