package service

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/rfsx0829/chatroom/server/mysql"
)

type Platform struct {
	OnlineList []*Person
	RoomList   []*Room
}

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

type Room struct {
	Id    int
	Name  string
	Pass  string
	IsPri bool

	OwnerId      int
	MessageQueue []*Message
	PersonList   []*Person
	PersonCount  int
	Token        string
}

type Message struct {
	Text string    `json:"text"`
	From int       `json:"from"`
	To   int       `json:"to"`
	Time time.Time `json:"time"`
}
