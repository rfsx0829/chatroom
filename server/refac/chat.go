package refac

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Platform platform
type Platform struct {
	RoomTable map[int]*Room
	UserTable map[int]*User
	ConnPool  map[int]*websocket.Conn

	errs   chan error
	cancel chan empty
}

// NewPlat create a platform
func NewPlat() *Platform {
	p := &Platform{
		RoomTable: make(map[int]*Room),
		UserTable: make(map[int]*User),
		ConnPool:  make(map[int]*websocket.Conn),

		errs:   make(chan error),
		cancel: make(chan empty),
	}
	go p.logerr()
	return p
}

// AddUser add user
func (p *Platform) AddUser(id int, name string, conn *websocket.Conn) {
	if _, ok := p.UserTable[id]; !ok {
		u := User{
			ID:          id,
			Name:        name,
			InWhichRoom: nil,
		}

		p.UserTable[id] = &u
		p.ConnPool[id] = conn
	}
}

// DelUser delete user
func (p *Platform) DelUser(id int) {
	delete(p.UserTable, id)
	delete(p.ConnPool, id)
}

// CreateRoom createroom
func (p *Platform) CreateRoom(name string) {
	id := p.getUnusedID()
	if id == -1 {
		return
	}

	r := Room{
		ID:     id,
		Name:   name,
		InRoom: make([]*User, 0, 5),
	}

	p.RoomTable[r.ID] = &r
}

func (p *Platform) getUnusedID() int {
	for i := 1; i < 300; i++ {
		if _, ok := p.RoomTable[i]; !ok {
			return i
		}
	}

	return -1
}

// DeleteRoom delete room
func (p *Platform) DeleteRoom(id int) {
	if r, ok := p.RoomTable[id]; ok {
		for _, e := range r.InRoom {
			e.InWhichRoom = nil
		}
		delete(p.RoomTable, id)
	}
}

// Enter room
func (p *Platform) Enter(uid, rid int) {
	if u, ok := p.UserTable[uid]; ok {
		if r, ok := p.RoomTable[rid]; ok {
			u.InWhichRoom = r
			r.InRoom = append(r.InRoom, u)
		}
	}
}

//Leave room
func (p *Platform) Leave(uid int) {
	if u, ok := p.UserTable[uid]; ok {
		if u.InWhichRoom != nil {
			u.InWhichRoom.removeUser(uid)
			u.InWhichRoom = nil
		}
	}
}

// SendMess sendmess
func (p *Platform) SendMess(from int, mes *Message) {
	if u, ok := p.UserTable[from]; ok {
		for _, e := range u.InWhichRoom.InRoom {
			p.send(e.ID, mes)
		}
	}
}

func (p *Platform) send(to int, mes *Message) {
	if con, ok := p.ConnPool[to]; ok {
		err := con.WriteJSON(mes)
		if err != nil {
			p.errs <- fmt.Errorf("[%d, %v]: %s", to, mes, err.Error())
		}
	}
}

func (p *Platform) logerr() {
	for {
		select {
		case err := <-p.errs:
			log.Println(err)
		case <-p.cancel:
			log.Println("PlatForm Log Canceld.")
			return
		}
	}
}

// User user
type User struct {
	ID          int
	Name        string
	InWhichRoom *Room
}

// Room room
type Room struct {
	ID     int
	Name   string
	InRoom []*User
}

func (r *Room) removeUser(uid int) {
	for i, e := range r.InRoom {
		if e.ID == uid {
			r.InRoom = append(r.InRoom[:i], r.InRoom[i+1:]...)
		}
	}
}

// Message mess
type Message struct {
	Content string
}

type empty struct{}
