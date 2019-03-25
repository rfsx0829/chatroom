package service

import (
	"github.com/gorilla/websocket"
	"github.com/rfsx0829/chatroom/server/mysql"
)

type Person struct {
	Info *mysql.Profile `json:"info"`

	Online     bool
	MessageBox []*Message
	Conn       *websocket.Conn

	/*
		-1 for not in room
	*/
	RoomId    int
	RoomToken string
}
