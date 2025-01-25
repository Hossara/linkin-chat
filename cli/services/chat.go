package services

import (
	"github.com/Hossara/linkin-chat/cli/pkg/api"
	"github.com/Hossara/linkin-chat/cli/types"
	"github.com/spf13/viper"
	"strings"
)

func GetAllChats() ([]types.ResponseChatRoom, error) {
	server := viper.GetString("server")
	a := api.NewApiHandler(server, 8080)

	response, err := api.Get[types.AllChatsResponse](a, "/chat", nil, api.GetAuthHeaders())

	if err != nil {
		return nil, err
	}

	return response.Chats, nil
}

func CreateNewChat(title string) (string, error) {
	server := viper.GetString("server")
	a := api.NewApiHandler(server, 8080)

	response, err := api.Post[types.CreateNewChatResponse](
		a, "/chat/"+strings.TrimSpace(strings.ToLower(title)),
		nil, nil, api.GetAuthHeaders(),
	)

	if err != nil {
		return "", err
	}

	return response.Code, nil
}
