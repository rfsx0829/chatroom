package controller

import "github.com/rfsx0829/chatroom/server/service"

type UserInfo struct {
	Uid   int    `json:"uid"`
	Name  string `json:"name"`
	Pass  string `json:"pass"`
	Email string `json:"email"`
}

type Room struct {
	Rid  int    `json:"rid"`
	Name string `json:"rname"`
	Pass string `json:"rpass"`
}

type FormData struct {
	Oper Op               `json:"oper"`
	User UserInfo         `json:"user"`
	Room Room             `json:"room"`
	Mes  *service.Message `json:"mes"`
}

type Response struct {
	Status int         `json:"status"`
	Text   string      `json:"text"`
	Extra  interface{} `json:"extra"`
}

/*
{
	"oper": 1,
	"user": {
		"uid": 123,
		"name": "name",
		"pass": "password",
		"Email": "email@email.com"
	},
	"room": {
		"rid": 1234,
		"rname": "rname",
		"rpass": "rpass"
	},
	"mes": "Hello World",
	"sendto": 124
}
*/
