package types

import (
	"time"

	"github.com/rfsx0829/chatroom/cmd/types/oper"
)

type FormData struct {
	Oper oper.Op  `json:"oper"`
	User UserInfo `json:"user"`
	Room Room     `json:"room"`
	Mes  Message  `json:"mes"`
}

type Response struct {
	Status int         `json:"status"`
	Text   string      `json:"text"`
	Extra  interface{} `json:"extra"`
}

type Room struct {
	Rid  int    `json:"rid"`
	Name string `json:"rname"`
	Pass string `json:"rpass"`
}

type UserInfo struct {
	Uid   int    `json:"uid"`
	Name  string `json:"name"`
	Pass  string `json:"pass"`
	Email string `json:"email"`
}

type Message struct {
	Text string    `json:"text"`
	From int       `json:"from"`
	To   int       `json:"to"`
	Time time.Time `json:"time"`
}
