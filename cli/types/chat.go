package types

import (
	"time"
)

type ResponseChatRoom struct {
	ID    uint   `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type AllChatsResponse struct {
	Chats []ResponseChatRoom `json:"chats"`
}

type CreateNewChatResponse struct {
	Code string `json:"code"`
}

type ChatRoomUser struct {
	Username string
	Role     uint
}

type Message struct {
	Sender    ChatRoomUser
	Content   string
	CreatedAt time.Time
}
