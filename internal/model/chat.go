package model

type Chat struct {
	Id       int64 `db:"id"`
	Username string
	Msg      string
}
