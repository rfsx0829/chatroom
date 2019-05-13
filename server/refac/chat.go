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
	UserTemp  []int

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
func (p *Platform) AddUser(id int, name string) {
	if _, ok := p.UserTable[id]; !ok {
		u := User{
			ID:          id,
			Name:        name,
			InWhichRoom: nil,
		}

		p.UserTable[id] = &u
	}

	if c, ok := p.ConnPool[id]; !ok || c == nil {
		p.UserTemp = append(p.UserTemp, id)
	}
}

// AddConn add conn
func (p *Platform) AddConn(conn *websocket.Conn) {
	if len(p.UserTemp) == 0 {
		return
	}

	id := p.UserTemp[0]
	p.UserTemp = p.UserTemp[1:]

	if u, ok := p.UserTable[id]; ok {
		p.ConnPool[id] = conn

		for _, v := range p.RoomTable {
			u.enterRoom(v)
		}
	}

	go p.routine(id, conn)
}

func (p *Platform) routine(id int, conn *websocket.Conn) {
	var x struct {
		Oper int    `json:"oper"`
		Mes  string `json:"mes"`
	}

	for {
		err := conn.ReadJSON(&x)
		if err != nil {
			log.Println("[routine] Error. Will Delete Conn.", err)
			p.UserTable[id].leaveRoom()
			p.ConnPool[id] = nil
			return
		}

		log.Println("[routine]", x)

		if x.Oper == 1 {
			p.SendMess(id, &Message{x.Oper, x.Mes})
		}
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
		inRoom: make([]*User, 0, 5),
		Nums:   0,
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
		for _, e := range r.inRoom {
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
			r.inRoom = append(r.inRoom, u)
			r.Nums++
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
		for _, e := range u.InWhichRoom.inRoom {
			p.send(e.ID, mes)
		}
	}
}

func (p *Platform) send(to int, mes *Message) {
	if con, ok := p.ConnPool[to]; ok {
		if err := con.WriteJSON(mes); err != nil {
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

// GetRoomList return rmlist
func (p *Platform) GetRoomList() []*Room {
	list := make([]*Room, 0, len(p.RoomTable))
	for _, v := range p.RoomTable {
		list = append(list, v)
	}
	return list
}

// User user
type User struct {
	ID          int
	Name        string
	InWhichRoom *Room
}

func (u *User) enterRoom(r *Room) {
	u.InWhichRoom = r
	r.inRoom = append(r.inRoom, u)
}

func (u *User) leaveRoom() {
	if u.InWhichRoom != nil {
		u.InWhichRoom.removeUser(u.ID)
		u.InWhichRoom = nil
	}
}

// Room room
type Room struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Nums   int    `json:"nums"`
	inRoom []*User
}

func (r *Room) removeUser(uid int) {
	for i, e := range r.inRoom {
		if e.ID == uid {
			r.inRoom = append(r.inRoom[:i], r.inRoom[i+1:]...)
			r.Nums--
		}
	}
}

// Message mess
type Message struct {
	Oper int    `json:"oper"`
	Mes  string `json:"mes"`
}

type empty struct{}
