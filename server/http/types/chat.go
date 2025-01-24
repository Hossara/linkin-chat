package types

type ResponseChatRoom struct {
	ID    uint
	Code  string
	Title string
}

type AllChatsResponse struct {
	Chats []ResponseChatRoom `json:"chats"`
}

type CreateNewChatResponse struct {
	Code string `json:"code"`
}
