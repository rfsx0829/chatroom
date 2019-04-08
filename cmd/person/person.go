package person

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rfsx0829/chatroom/cmd/conn"
	"github.com/rfsx0829/chatroom/cmd/types"
	"github.com/rfsx0829/chatroom/cmd/types/oper"
)

// Person represent a person.
type Person struct {
	Info       *types.UserInfo
	MessageBox []*types.Message
	Conn       *conn.Conn
	FormData   *types.FormData
	RoomID     int
	RoomToken  string
}

// AddEmail just do the thing like its name.
func (p *Person) AddEmail(email string) error {
	p.FormData.Oper = oper.AddEmail
	p.FormData.User.Email = email

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	return nil
}

// CreateRoom with room name and password.
func (p *Person) CreateRoom(name, pass string) error {
	p.FormData.Oper = oper.Create
	p.FormData.Room.Name = name
	p.FormData.Room.Pass = pass

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	if res.Status == http.StatusOK {
		rid := res.Extra.(map[string]interface{})["rid"].(float64)
		p.RoomID = int(rid)
	}
	return nil
}

// EnterRoom let the person enter the room use password.
func (p *Person) EnterRoom(rid int, pass string) error {
	p.FormData.Oper = oper.Enter
	p.FormData.Room.Rid = rid
	p.FormData.Room.Pass = pass

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	return nil
}

// LeaveRoom just leave room.
func (p *Person) LeaveRoom() error {
	p.FormData.Oper = oper.Leave

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	return nil
}

// SendMessage send a message to room.
func (p *Person) SendMessage(content string) error {
	p.FormData.Oper = oper.SendMes
	p.FormData.Mes.From = p.Info.UID
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

// SendToBox send a message to single person.
func (p *Person) SendToBox(content string, sendTo int) error {
	p.FormData.Oper = oper.SendBox
	p.FormData.Mes.From = p.Info.UID
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

// GetRoomList get the room list
func (p *Person) GetRoomList() error {
	p.FormData.Oper = oper.GetRoomList

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	// TODO:
	return nil
}

// GetPersonsInRoom get the list of online person in room.
func (p *Person) GetPersonsInRoom() error {
	p.FormData.Oper = oper.GetPersonsInRoom

	res, err := p.sendReq()
	if err != nil {
		return err
	}

	fmt.Println(res.Text)
	// TODO:
	return nil
}

// GoOffline close the conn of a person.
func (p *Person) GoOffline() error {
	p.FormData.Oper = oper.Close
	return p.Conn.Close(p.FormData)
}

func (p *Person) sendReq() (*types.Response, error) {
	return p.Conn.WriteAndRead(p.FormData)
}
