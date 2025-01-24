package services

import (
	"context"
	"github.com/Hossara/linkin-chat/app"
	chatPort "github.com/Hossara/linkin-chat/internal/chat/port"
	userDomain "github.com/Hossara/linkin-chat/internal/user/domain"
	"github.com/Hossara/linkin-chat/server/http/helpers"
	"github.com/Hossara/linkin-chat/server/http/types"
)

type ChatService struct {
	svc chatPort.Service
}

func NewChatService(
	svc chatPort.Service,
) *ChatService {
	return &ChatService{
		svc: svc,
	}
}

func ChatServiceGetter(appContainer app.App) helpers.ServiceGetter[*ChatService] {
	return func(ctx context.Context) *ChatService {
		return NewChatService(
			appContainer.ChatService(),
		)
	}
}

func (as *ChatService) GetAllChats(c context.Context, userID uint) (*types.AllChatsResponse, error) {
	rooms, err := as.svc.GetUserChatRooms(c, userDomain.UserID(userID))

	if err != nil {
		return nil, err
	}

	return &types.AllChatsResponse{
		Chats: rooms,
	}, nil
}
