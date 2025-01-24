package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"

	"github.com/Hossara/linkin-chat/cli/pages"
)

var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join the chatroom",
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")

		pages.LoginPage(username, password, server)

		log.Println(viper.GetString("login.token"))
	},
}

func SetJoinCommand(cmd *cobra.Command) {
	cmd.AddCommand(joinCmd)

	joinCmd.Flags().StringP("server", "s", "localhost", "Server address")
}
