package services

import (
	"github.com/Hossara/linkin-chat/cli/pkg/api"
	"github.com/Hossara/linkin-chat/cli/types"
	"strings"
)

func Login(username, password, server string) (string, error) {
	loginBody := &api.RequestBody{
		"username": strings.TrimSpace(username),
		"password": strings.TrimSpace(password),
	}

	a := api.NewApiHandler(server, 8080)

	response, err := api.Post[types.LoginResponse](a, "/auth/login", loginBody, nil)

	if err != nil {
		return "", err
	}

	return response.Token, nil
}

func Register(username, password, firstName, lastName, server string) (string, error) {
	registerBody := &api.RequestBody{
		"username":   strings.TrimSpace(username),
		"password":   strings.TrimSpace(password),
		"first_name": strings.TrimSpace(firstName),
		"last_name":  strings.TrimSpace(lastName),
	}

	a := api.NewApiHandler(server, 8080)

	response, err := api.Post[types.LoginResponse](a, "/auth/register", registerBody, nil)

	if err != nil {
		return "", err
	}

	return response.Token, nil
}
