package services

import (
	"context"
	"github.com/Hossara/linkin-chat/app"
	"github.com/Hossara/linkin-chat/internal/chat/domain"
	"github.com/Hossara/linkin-chat/pkg/utils"
	"github.com/Hossara/linkin-chat/server/http/helpers"
	"github.com/Hossara/linkin-chat/server/http/mapper"
	"github.com/Hossara/linkin-chat/server/http/types"

	chatService "github.com/Hossara/linkin-chat/internal/chat"
	chatPort "github.com/Hossara/linkin-chat/internal/chat/port"
	userDomain "github.com/Hossara/linkin-chat/internal/user/domain"
)

var (
	ErrMaximumChatReached = chatService.ErrMaximumChatReached
	ErrInvalidChatInfo    = chatService.ErrInvalidChatInfo
	ErrInvalidUserID      = chatService.ErrInvalidUserID
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

	return mapper.ToResponseChatRoom(rooms), nil
}

func (as *ChatService) CreateNewChat(c context.Context, userID uint, title string) (*types.CreateNewChatResponse, error) {
	title = utils.NormalizeString(title)

	code, err := as.svc.CreateChatRoom(c, domain.ChatRoom{
		Title:   title,
		Users:   []userDomain.User{},
		OwnerID: userDomain.UserID(userID),
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateNewChatResponse{
		Code: string(code),
	}, nil
}
