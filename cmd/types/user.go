package types

import (
	"fmt"
	"net/http"
	"time"
)

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

type Person struct {
	Info       *UserInfo
	MessageBox []*Message
	Conn       *Conn
	FormData   *FormData
	RoomId     int
	RoomToken  string
}

func (p *Person) AddEmail(email string) error {
	p.FormData.Oper = AddEmail
	p.FormData.User.Email = email

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	return nil
}

func (p *Person) CreateRoom(name, pass string) error {
	p.FormData.Oper = Create
	p.FormData.Room.Name = name
	p.FormData.Room.Pass = pass

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	if res.Status == http.StatusOK {
		rid := res.Extra.(map[string]interface{})["rid"].(float64)
		p.RoomId = int(rid)
	}
	return nil
}

func (p *Person) EnterRoom(rid int, pass string) error {
	p.FormData.Oper = Enter
	p.FormData.Room.Rid = rid
	p.FormData.Room.Pass = pass

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	return nil
}

func (p *Person) LeaveRoom() error {
	p.FormData.Oper = Leave

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	return nil
}

func (p *Person) SendMessage(content string) error {
	p.FormData.Oper = SendMes
	p.FormData.Mes.From = p.Info.Uid
	p.FormData.Mes.Text = content
	p.FormData.Mes.To = -1
	p.FormData.Mes.Time = time.Now()

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	return nil
}

func (p *Person) SendToBox(content string, sendTo int) error {
	p.FormData.Oper = SendBox
	p.FormData.Mes.From = p.Info.Uid
	p.FormData.Mes.Text = content
	p.FormData.Mes.To = sendTo
	p.FormData.Mes.Time = time.Now()

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	return nil
}

func (p *Person) GetRoomList() error {
	p.FormData.Oper = GetRoomList

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	// TODO:
	return nil
}

func (p *Person) GetPersonsInRoom() error {
	p.FormData.Oper = GetPersonsInRoom

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	// TODO:
	return nil
}

func (p *Person) GoOffline() error {
	p.FormData.Oper = Close
	return p.Conn.Close(p.FormData)
}

func (p *Person) sendReq() (*Response, error) {
	return p.Conn.WriteAndRead(p.FormData)
}
