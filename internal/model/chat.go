package model

type Chat struct {
	Id       int64
	Username string
	Msg      string
}

type InMessage struct {
	UserID   int    `json:"id"`
	UserName string `json:"user_name"`
	ChatID   string `json:"chat_id"`
	Text     string `json:"text"`
}

type OutMessage struct {
	From string
	Text string
}
