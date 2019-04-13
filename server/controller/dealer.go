package controller

import (
	"errors"

	"github.com/gorilla/websocket"
	"github.com/rfsx0829/chatroom/server/service"
)

type Dealer func(*FormData, *Response) error
type Op uint32

const (
	DefaultOper Op = iota
	SignUp
	SignIn
	AddEmail
	Create
	Enter
	Leave
	SendMes
	SendBox
	GetRoomList
	GetPersonsInRoom
	Close
)

var ops = map[Op]Dealer{
	DefaultOper: DefaultOperDealer,
	SignUp:      SignUpDealer,
	AddEmail:    AddEmailDealer,
	Create:      CreateDealer,
	Enter:       EnterDealer,
	Leave:       LeaveDealer,
	SendMes:     SendMesDealer,
	SendBox:     SendBoxDealer,
	GetRoomList: GetRoomListDealer,
}

var strs = map[Op]string{
	DefaultOper:      "DefaultOper",
	SignUp:           "SignUp",
	SignIn:           "SignIn",
	AddEmail:         "AddEmail",
	Create:           "Create",
	Enter:            "Enter",
	Leave:            "Leave",
	SendMes:          "SendMes",
	SendBox:          "SendBox",
	GetRoomList:      "GetRoomList",
	GetPersonsInRoom: "GetPersonsInRoom",
	Close:            "Close",
}

func (o Op) String() string {
	if str, found := strs[o]; found {
		return str
	}
	return "Unknown"
}

func DefaultOperDealer(x *FormData, res *Response) error {
	return errors.New("Unknown Operation, Please Check !")
}

func SignInDealer(con *websocket.Conn, x *FormData, res *Response) error {
	var (
		uid int
		err error
	)

	if x.User.Name != "" {
		uid, err = service.Client.SignInWithName(con, x.User.Name, x.User.Pass)
	} else {
		uid, err = service.Client.SignInWithEmail(con, x.User.Email, x.User.Pass)
	}

	if err != nil {
		return err
	}

	res.Text = "SignIn Success !"
	res.Extra = struct {
		Uid int `json:"uid"`
	}{uid}
	return nil
}

func SignUpDealer(x *FormData, res *Response) error {
	err := service.Client.SignUp(x.User.Name, x.User.Pass)
	if err != nil {
		return err
	}

	res.Text = "SignUp Success !"
	res.Extra = nil
	return nil
}

func AddEmailDealer(x *FormData, res *Response) error {
	err := service.Client.AddEmail(x.User.Uid, x.User.Email)
	if err != nil {
		return err
	}

	res.Text = "AddEmail Success !"
	res.Extra = nil
	return nil
}

func CreateDealer(x *FormData, res *Response) error {
	rid := service.Client.CreateRoom(x.Room.Name, x.Room.Pass, false, x.User.Uid)

	res.Text = "Create Room Success !"
	res.Extra = struct {
		Rid int `json:"rid"`
	}{rid}
	return nil
}

func EnterDealer(x *FormData, res *Response) error {
	err := service.Client.EnterRoom(x.User.Uid, x.Room.Rid, x.Room.Pass)
	if err != nil {
		return err
	}

	res.Text = "Enter Room Success !"
	res.Extra = nil
	return nil
}

func LeaveDealer(x *FormData, res *Response) error {
	err := service.Client.LeaveRoom(x.User.Uid)
	if err != nil {
		return err
	}

	res.Text = "LeaveRoom Success !"
	res.Extra = nil
	return nil
}

func SendMesDealer(x *FormData, res *Response) error {
	err := service.Client.SendMes(x.User.Uid, x.Mes)
	if err != nil {
		return err
	}

	res.Text = "Send Message Success !"
	res.Extra = nil
	return nil
}

func SendBoxDealer(x *FormData, res *Response) error {
	err := service.Client.SendToBox(x.User.Uid, x.Mes)
	if err != nil {
		return err
	}

	res.Text = "Send To Box Success !"
	res.Extra = nil
	return nil
}

func GetRoomListDealer(x *FormData, res *Response) error {
	var list = []*Room{}
	for _, e := range service.Client.RoomList {
		rm := Room{
			Rid:  e.Id,
			Name: e.Name,
		}
		list = append(list, &rm)
	}

	res.Extra = list
	res.Text = "GetRoomList Success !"
	return nil
}

func GetPersonsInRoomDealer(x *FormData, res *Response) error {
	ur := service.Client.GetUserById(x.User.Uid)
	if ur == nil {
		return service.InvalidUid
	}

	if ur.RoomId == -1 {
		return errors.New("Not In A Room !")
	}

	rm := service.Client.GetRoomById(ur.RoomId)
	if rm == nil {
		return service.InvalidRid
	}

	list := []*UserInfo{}
	for _, e := range rm.PersonList {
		list = append(list, &UserInfo{
			Uid:   e.Info.Uid,
			Name:  e.Info.Name,
			Email: e.Info.Email,
		})
	}

	res.Extra = list
	res.Text = "GetPersonsInRoom Success !"
	return nil
}
