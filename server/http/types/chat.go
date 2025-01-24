package types

import chatDomain "github.com/Hossara/linkin-chat/internal/chat/domain"

type AllChatsResponse struct {
	Chats []*chatDomain.ChatRoom `json:"chats"`
}
