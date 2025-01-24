package commands

import (
	"github.com/Hossara/linkin-chat/cli/pages"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Create account in linkin chat",
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")

		pages.RegisterPage(server)
	},
}

func SetRegisterCommand(cmd *cobra.Command) {
	cmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("server", "s", "localhost", "Server address")
}
