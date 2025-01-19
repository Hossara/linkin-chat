package commands

import (
	"github.com/spf13/cobra"
	"log"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login or create account",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic to join the chatroom
		log.Println("Joining the chatroom...")
		// Connect to server logic here (e.g., via NATS)
	},
}

func SetLoginCommand(cmd *cobra.Command) {
	cmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "u", "guest", "LinkinChat account username")
	loginCmd.Flags().StringP("password", "p", "123456", "LinkinChat account password")
}
