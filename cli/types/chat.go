package types

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
