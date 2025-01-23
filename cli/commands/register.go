package commands

import (
	"github.com/Hossara/linkin-chat/cli/pages"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Create account in linkin chat",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		server, _ := cmd.Flags().GetString("server")

		pages.RegisterPage(username, password, server)
	},
}

func SetRegisterCommand(cmd *cobra.Command) {
	cmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("username", "u", "", "Account username")
	registerCmd.Flags().StringP("password", "p", "", "Account password")
	registerCmd.Flags().StringP("first_name", "f", "", "User first name")
	registerCmd.Flags().StringP("last_name", "l", "", "User last name")
	registerCmd.Flags().StringP("server", "s", "localhost", "Server address")
}
