package mapper

import (
	chatDomain "github.com/Hossara/linkin-chat/internal/chat/domain"
	"github.com/Hossara/linkin-chat/pkg/utils"
	"github.com/Hossara/linkin-chat/server/http/types"
)

func ToResponseChatRoom(chats []*chatDomain.ChatRoom) *types.AllChatsResponse {
	return &types.AllChatsResponse{
		Chats: utils.Map(chats, func(t *chatDomain.ChatRoom) types.ResponseChatRoom {
			return types.ResponseChatRoom{
				ID:    uint(t.ID),
				Title: t.Title,
				Code:  string(t.Code),
			}
		}),
	}
}
