package commands

import (
	"github.com/spf13/cobra"
	"log"
)

var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join the chatroom",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic to join the chatroom
		log.Println("Joining the chatroom...")
		// Connect to server logic here (e.g., via NATS)
	},
}

func SetJoinCommand(cmd *cobra.Command) {
	cmd.AddCommand(joinCmd)

	joinCmd.Flags().StringP("server", "s", "localhost:4222", "NATS server address")
}
