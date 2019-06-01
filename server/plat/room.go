package plat

import (
	"encoding/json"
	"errors"
)

// Room room
type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Nums int    `json:"nums"`

	inRoom   []*User
	messages []*wsMes
	pass     string
}

func (r *Room) removeUser(uid int) {
	for i, e := range r.inRoom {
		if e.ID == uid {
			r.inRoom = append(r.inRoom[:i], r.inRoom[i+1:]...)
			r.Nums--
		}
	}
}

func (r *Room) addUser(u *User) {
	for _, e := range r.inRoom {
		if e.ID == u.ID {
			return
		}
	}
	r.inRoom = append(r.inRoom, u)
	r.Nums++
}

// GetRoomList return rmlist
func (p *Platform) GetRoomList() ([]byte, error) {
	list := make([]*Room, 0, len(p.RoomTable))
	for _, v := range p.RoomTable {
		list = append(list, v)
	}

	return json.Marshal(list)
}

// GetRoomMessages return room messages
func (p *Platform) GetRoomMessages(rid int) ([]byte, error) {
	if r, ok := p.RoomTable[rid]; ok {
		return json.Marshal(r.messages)
	}
	return nil, errors.New("No Such Room")
}

// CreateRoom create room
func (p *Platform) CreateRoom(name, pass string) (id int, err error) {
	id = p.getUnusedID()
	if id == -1 {
		return 0, errors.New("Too Many Rooms")
	}

	r := Room{
		ID:       id,
		Name:     name,
		pass:     pass,
		messages: make([]*wsMes, 0, 5),
		inRoom:   make([]*User, 0, 5),
		Nums:     0,
	}

	p.RoomTable[r.ID] = &r
	return id, nil
}

// DeleteRoom delete room
func (p *Platform) DeleteRoom(id int) {
	if r, ok := p.RoomTable[id]; ok {
		for _, e := range r.inRoom {
			r.removeUser(e.ID)
		}
		delete(p.RoomTable, id)
	}
}

// Enter room
func (p *Platform) Enter(uid, rid int, pass string) error {
	if u, ok := p.UserTable[uid]; ok {
		if r, ok := p.RoomTable[rid]; ok {
			if !checkPassword(r.pass, pass) {
				return errors.New("Invalid password")
			}

			if u.inWhichRoom != nil {
				u.inWhichRoom.removeUser(u.ID)
			}
			r.addUser(u)
			u.inWhichRoom = r
			return nil
		}
		return errors.New("No Such Room")
	}
	return errors.New("No Such User")
}

// Leave room
func (p *Platform) Leave(uid int) {
	if u, ok := p.UserTable[uid]; ok {
		if u.inWhichRoom != nil {
			u.inWhichRoom.removeUser(uid)
			u.inWhichRoom = nil
		}
	}
}
