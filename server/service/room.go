package service

import "net/http"

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

func (r *Room) broadCastMes(mes *Message) (err error) {
	r.MessageQueue = append(r.MessageQueue, mes)
	res := Response{
		Status: http.StatusOK,
		Oper:   7, // 7 for SendMes
		Text:   "Receive Message",
		Extra:  mes,
	}

	for _, e := range r.PersonList {
		err = e.Conn.WriteJSON(res)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Platform) CreateRoom(name, pass string, isPri bool, ownerId int) int {
	return p.createRoom(name, pass, isPri, ownerId)
}

func (p *Platform) createRoom(name, pass string, isPri bool, ownerId int) int {
	r := &Room{
		Id:    p.getUnUsedId(),
		Name:  name,
		Pass:  pass,
		IsPri: isPri,

		OwnerId:      ownerId,
		MessageQueue: make([]*Message, 0, 10),
		PersonList:   make([]*Person, 0, 5),
		PersonCount:  0,
		Token:        "",
	}

	p.RoomList = append(p.RoomList, r)
	p.EnterRoom(r.Id, ownerId, pass)
	return r.Id
}

func (p *Platform) EnterRoom(uid, rid int, pass string) error {
	ur, rm := p.GetUserById(uid), p.GetRoomById(rid)
	if ur == nil {
		return InvalidUid
	}
	if rm == nil {
		return InvalidRid
	}

	if rm.Pass != pass {
		return InvalidPass
	}

	ur.RoomId = rm.Id
	ur.RoomToken = rm.Token

	rm.PersonCount++
	rm.PersonList = append(rm.PersonList, ur)
	return nil
}

func (p *Platform) LeaveRoom(uid int) error {
	ur := p.GetUserById(uid)
	if ur == nil {
		return InvalidUid
	}

	rm := p.GetRoomById(ur.RoomId)
	if rm == nil {
		return InvalidRid
	}

	rm.PersonCount--
	if rm.PersonCount == 0 {
		p.DeleteRoom(rm)
		return nil
	}

	for i, e := range rm.PersonList {
		if e.Info.Uid == uid {
			if i == len(rm.PersonList)-1 {
				rm.PersonList = rm.PersonList[0:i]
				break
			}
			rm.PersonList = append(rm.PersonList[0:i], rm.PersonList[i+1:]...)
			break
		}
	}

	ur.RoomId = -1
	ur.RoomToken = ""
	return nil
}

func (p *Platform) DeleteRoom(rm *Room) {
	for i, e := range p.RoomList {
		if e.Id == rm.Id {
			if i == len(p.RoomList)-1 {
				p.RoomList = p.RoomList[0:i]
				break
			}
			p.RoomList = append(p.RoomList[0:i], p.RoomList[i+1:]...)
			break
		}
	}
}
