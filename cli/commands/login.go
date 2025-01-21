package commands

import (
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login or create account",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		if username == "" || password == "" {

		}
	},
}

func SetLoginCommand(cmd *cobra.Command) {
	cmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "u", "", "LinkinChat account username")
	loginCmd.Flags().StringP("password", "p", "", "LinkinChat account password")
	loginCmd.Flags().StringP("server", "s", "localhost", "Server address")
}
