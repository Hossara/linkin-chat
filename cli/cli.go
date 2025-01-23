package cli

import (
	"fmt"
	"github.com/Hossara/linkin-chat/cli/constants"
	"github.com/spf13/cobra"
	"os"

	"github.com/Hossara/linkin-chat/cli/commands"
)

var rootCmd = &cobra.Command{
	Use:   "linkin-chat",
	Short: "Linkin Chat Nats Based Chatroom CLI Application",
	Long:  "A Nats-based chatroom application where users can join, send messages, and view active users.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(constants.Logo)
		fmt.Print("\n\n")
		fmt.Println("Welcome to the Linkin Chat CLI. Use --help for available commands.")
	},
}

func Execute() {
	commands.SetJoinCommand(rootCmd)
	commands.SetRegisterCommand(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
