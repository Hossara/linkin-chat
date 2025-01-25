package commands

import (
	"github.com/Hossara/linkin-chat/cli/pages"
	"github.com/Hossara/linkin-chat/cli/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join the chatroom",
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")

		viper.Set("server", server)
		username := viper.GetString("login.username")
		password := viper.GetString("login.password")

		if username == "" || password == "" {
			pages.LoginPage()
		} else {
			token, err := services.Login(username, password)

			viper.Set("login.token", token)

			if err != nil {
				viper.Set("login", nil)
				err := viper.WriteConfig()

				if err != nil {
					log.Fatalf("Error while writing config: %v", err)
					return
				}

				pages.LoginPage()
			}
		}

		token := viper.GetString("login.token")

		if token == "" {
			log.Fatalf("Token is empty! Something wrong with configuration file!")
		}

		pages.HomePage()

	},
}

func SetJoinCommand(cmd *cobra.Command) {
	cmd.AddCommand(joinCmd)

	joinCmd.Flags().StringP("server", "s", "localhost", "Server address")
}
