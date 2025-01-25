package services

import (
	"github.com/Hossara/linkin-chat/cli/pkg/api"
	"github.com/Hossara/linkin-chat/cli/types"
	"github.com/spf13/viper"
	"strings"
)

func Login(username, password string) (string, error) {
	loginBody := &api.RequestBody{
		"username": strings.TrimSpace(username),
		"password": strings.TrimSpace(password),
	}

	server := viper.GetString("server")
	a := api.NewApiHandler(server, 8080)

	response, err := api.Post[types.LoginResponse](a, "/auth/login", loginBody, nil, nil)

	if err != nil {
		return "", err
	}

	return response.Token, nil
}

func Register(username, password, firstName, lastName string) (string, error) {
	registerBody := &api.RequestBody{
		"username":   strings.TrimSpace(username),
		"password":   strings.TrimSpace(password),
		"first_name": strings.TrimSpace(firstName),
		"last_name":  strings.TrimSpace(lastName),
	}

	server := viper.GetString("server")
	a := api.NewApiHandler(server, 8080)

	response, err := api.Post[types.LoginResponse](a, "/auth/register", registerBody, nil, nil)

	if err != nil {
		return "", err
	}

	return response.Token, nil
}
