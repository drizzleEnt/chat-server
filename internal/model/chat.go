package model

type Chat struct {
	Id   int64    `db:""id`
	Name []string `db:"username"`
	Msg  string   `db:"message"`
}
