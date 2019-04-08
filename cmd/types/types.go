package types

import (
	"time"

	"github.com/rfsx0829/chatroom/cmd/types/oper"
)

// FormData is a data pack.
type FormData struct {
	Oper oper.Op  `json:"oper"`
	User UserInfo `json:"user"`
	Room RoomInfo `json:"room"`
	Mes  Message  `json:"mes"`
}

// Response represent a response from server.
type Response struct {
	Status int         `json:"status"`
	Text   string      `json:"text"`
	Extra  interface{} `json:"extra"`
}

// RoomInfo just maintain the information of a room.
type RoomInfo struct {
	Rid  int    `json:"rid"`
	Name string `json:"rname"`
	Pass string `json:"rpass"`
}

// UserInfo just maintain the basic information of a room.
type UserInfo struct {
	UID   int    `json:"uid"`
	Name  string `json:"name"`
	Pass  string `json:"pass"`
	Email string `json:"email"`
}

// Message represent a message that person sent.
type Message struct {
	Text string    `json:"text"`
	From int       `json:"from"`
	To   int       `json:"to"`
	Time time.Time `json:"time"`
}
