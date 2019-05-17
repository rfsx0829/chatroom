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
	cancel chan struct{}
}

// NewPlat create a platform
func NewPlat() *Platform {
	p := &Platform{
		RoomTable: make(map[int]*Room),
		UserTable: make(map[int]*User),
		ConnPool:  make(map[int]*websocket.Conn),

		errs:   make(chan error),
		cancel: make(chan struct{}),
	}
	// go p.logerr()

	p.RoomTable[1] = &Room{
		ID:     1,
		Name:   "DefaultRoom",
		Nums:   0,
		inRoom: make([]*User, 0, 2),
	}

	return p
}

// AddUser add user
func (p *Platform) AddUser(id int, name string) {
	if _, ok := p.UserTable[id]; !ok {
		u := User{
			ID:          id,
			Name:        name,
			InWhichRoom: p.RoomTable[1],
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
		u.InWhichRoom = p.RoomTable[1]
	}

	go p.routine(id, conn)
}

// DelUser delete user
func (p *Platform) DelUser(id int) {
	delete(p.UserTable, id)
	delete(p.ConnPool, id)
}

// GetRoomList return rmlist
func (p *Platform) GetRoomList() []*Room {
	list := make([]*Room, 0, len(p.RoomTable))
	for _, v := range p.RoomTable {
		list = append(list, v)
	}
	return list
}

func (p *Platform) routine(id int, conn *websocket.Conn) {
	var x wsMes

	for {
		err := conn.ReadJSON(&x)
		if err != nil {
			log.Println("[routine] Error. Will Delete Conn.", err)
			p.UserTable[id].InWhichRoom = p.RoomTable[1]
			p.ConnPool[id] = nil
			return
		}

		log.Println("[routine]", x)

		p.dealer(id, &x)
	}
}

func (p *Platform) dealer(id int, mes *wsMes) {
	switch mes.Oper {
	case createRm:
		p.createRoom(mes.Mes)
	case deleteRm:
		p.deleteRoom(mes.ID)
	case enterRm:
		p.enter(id, mes.ID)
	case leaveRm:
		p.leave(id)
	case sendMes:
		p.sendMess(id, mes)
	}
}

func (p *Platform) createRoom(name string) {
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
	for i := 2; i < 300; i++ {
		if _, ok := p.RoomTable[i]; !ok {
			return i
		}
	}

	return -1
}

func (p *Platform) deleteRoom(id int) {
	if r, ok := p.RoomTable[id]; ok {
		for _, e := range r.inRoom {
			e.InWhichRoom = p.RoomTable[1]
		}
		delete(p.RoomTable, id)
	}
}

func (p *Platform) enter(uid, rid int) {
	if u, ok := p.UserTable[uid]; ok {
		if r, ok := p.RoomTable[rid]; ok {
			u.InWhichRoom = r
			r.inRoom = append(r.inRoom, u)
			r.Nums++
		}
	}
}

func (p *Platform) leave(uid int) {
	if u, ok := p.UserTable[uid]; ok {
		u.InWhichRoom = p.RoomTable[1]
	}
}

func (p *Platform) sendMess(from int, mes *wsMes) {
	if u, ok := p.UserTable[from]; ok {
		for _, e := range u.InWhichRoom.inRoom {
			p.send(e.ID, mes)
		}
	}
}

func (p *Platform) send(to int, mes *wsMes) {
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

type wsMes struct {
	Oper oper `json:"oper"`

	ID  int    `json:"id"`
	Mes string `json:"mes"`
}

type oper uint

const (
	createRm oper = iota
	deleteRm
	enterRm
	leaveRm
	sendMes
)
