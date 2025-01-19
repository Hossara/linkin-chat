package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/Hossara/linkin-chat/cli/commands"
)

var rootCmd = &cobra.Command{
	Use:   "linkin-chat",
	Short: "Linkin Chat Nats Based Chatroom CLI Application",
	Long:  "A Nats-based chatroom application where users can join, send messages, and view active users.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n   __ _       _    _           ___ _           _   \n  / /(_)_ __ | | _(_)_ __     / __\\ |__   __ _| |_ \n / / | | '_ \\| |/ / | '_ \\   / /  | '_ \\ / _` | __|\n/ /__| | | | |   <| | | | | / /___| | | | (_| | |_ \n\\____/_|_| |_|_|\\_\\_|_| |_| \\____/|_| |_|\\__,_|\\__|")
		fmt.Print("\n\n")
		fmt.Println("Welcome to the Linkin Chat CLI. Use --help for available commands.")
	},
}

func Execute() {
	commands.SetJoinCommand(rootCmd)
	commands.SetLoginCommand(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
