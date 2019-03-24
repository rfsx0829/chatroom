package controller

import (
	"errors"

	"github.com/rfsx0829/chatroom/server/service"
)

type Dealer func(*FormData, *Response) error

var ops = map[Op]Dealer{
	DefaultOper: DefaultOperDealer,
	SignUp:      SignUpDealer,
	AddEmail:    AddEmailDealer,
	Create:      CreateDealer,
	Enter:       EnterDealer,
	Leave:       LeaveDealer,
	SendMes:     SendMesDealer,
	SendBox:     SendBoxDealer,
}

func DefaultOperDealer(x *FormData, res *Response) error {
	return errors.New("Unknown Operation, Please Check !")
}

func SignUpDealer(x *FormData, res *Response) error {
	err := service.Client.SignUp(x.User.Name, x.User.Pass)
	if err != nil {
		return err
	}

	res.Text = "SignUp Success !"
	return nil
}

func AddEmailDealer(x *FormData, res *Response) error {
	err := service.Client.AddEmail(x.User.Uid, x.User.Email)
	if err != nil {
		return err
	}

	res.Text = "AddEmail Success !"
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
	return nil
}

func LeaveDealer(x *FormData, res *Response) error {
	err := service.Client.LeaveRoom(x.User.Uid)
	if err != nil {
		return err
	}

	res.Text = "LeaveRoom Success !"
	return nil
}

func SendMesDealer(x *FormData, res *Response) error {
	err := service.Client.SendMes(x.User.Uid, x.Mes)
	if err != nil {
		return err
	}

	res.Text = "Send Message Success !"
	return nil
}

func SendBoxDealer(x *FormData, res *Response) error {
	err := service.Client.SendToBox(x.User.Uid, x.SendTo, x.Mes)
	if err != nil {
		return err
	}

	res.Text = "Send To Box Success !"
	return nil
}
