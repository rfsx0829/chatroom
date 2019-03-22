package service

import (
	"github.com/gorilla/websocket"
)

var (
	Client   *Platform
	upgrader = websocket.Upgrader{}
)

func init() {
	Client = NewPlatform()
}

func NewPlatform() *Platform {
	if Client != nil {
		return nil
	}
	return &Platform{
		OnlineList: make([]*Person, 0, 10),
		RoomList:   make([]*Room, 0, 2),
	}
}

func (p *Platform) CreateRoom(name, pass string, isPri bool, ownerId int) *Room {
	return p.createRoom(name, pass, isPri, ownerId)
}

func (p *Platform) createRoom(name, pass string, isPri bool, ownerId int) *Room {
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
	return r
}

func (p *Platform) getUnUsedId() int {
	var (
		id = len(p.RoomList)
		ok = false
	)

	for !ok {
		for _, e := range p.RoomList {
			if e.Id == id {
				id++
				break
			}
		}

		ok = true
	}

	return id
}

func (p *Platform) EnterRoom(uid, rid int, pass string) error {
	ur, rm := p.getUserById(uid), p.getRoomById(rid)
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
	ur := p.getUserById(uid)
	if ur == nil {
		return InvalidUid
	}

	rm := p.getRoomById(ur.RoomId)
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

func (p *Platform) getUserById(uid int) *Person {
	for _, e := range p.OnlineList {
		if e.Info.Uid == uid {
			return e
		}
	}
	return nil
}

func (p *Platform) getRoomById(rid int) *Room {
	for _, e := range p.RoomList {
		if e.Id == rid {
			return e
		}
	}
	return nil
}
