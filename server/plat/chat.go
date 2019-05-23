package plat

import (
	"log"

	"github.com/gorilla/websocket"
)

// Platform platform
type Platform struct {
	RoomTable map[int]*Room
	UserTable map[int]*User
	ConnPool  map[int]*websocket.Conn
	UserTemp  []int
}

// New create a platform
func New() *Platform {
	p := &Platform{
		RoomTable: make(map[int]*Room),
		UserTable: make(map[int]*User),
		ConnPool:  make(map[int]*websocket.Conn),
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
			inWhichRoom: p.RoomTable[1],
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
		u.inWhichRoom = p.RoomTable[1]
		p.RoomTable[1].inRoom = append(p.RoomTable[1].inRoom, u)
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
	var x Message

	for {
		err := conn.ReadJSON(&x)
		if err != nil {
			log.Println("[routine] Error. Will Delete Conn.", err)
			p.UserTable[id].inWhichRoom.removeUser(id)
			p.UserTable[id].inWhichRoom = p.RoomTable[1]
			p.ConnPool[id] = nil
			return
		}

		log.Println("[routine]", x)

		p.broadCastMes(&x)
	}
}

// CreateRoom create room
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
	for i := 2; i < 300; i++ {
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
			e.inWhichRoom = p.RoomTable[1]
		}
		delete(p.RoomTable, id)
	}
}

// Enter room
func (p *Platform) Enter(uid, rid int) {
	if u, ok := p.UserTable[uid]; ok {
		if r, ok := p.RoomTable[rid]; ok {
			u.inWhichRoom = r
			r.inRoom = append(r.inRoom, u)
			r.Nums++
		}
	}
}

// Leave room
func (p *Platform) Leave(uid int) {
	if u, ok := p.UserTable[uid]; ok {
		u.inWhichRoom = p.RoomTable[1]
	}
}

func (p *Platform) broadCastMes(mes *Message) {
	if u, ok := p.UserTable[mes.User.ID]; ok {
		for _, e := range u.inWhichRoom.inRoom {
			p.sendToConn(e.ID, mes)
		}
	}
}

func (p *Platform) sendToConn(to int, mes *Message) {
	if con, ok := p.ConnPool[to]; ok {
		if err := con.WriteJSON(mes); err != nil {
			log.Printf("[%d, %v]: %s", to, mes, err.Error())
		}
	}
}

// User user
type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	inWhichRoom *Room
}

func (u *User) enterRoom(r *Room) {
	u.inWhichRoom = r
	r.inRoom = append(r.inRoom, u)
}

// Room room
type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Nums int    `json:"nums"`

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

// Message message
type Message struct {
	ID   int    `json:"id"`
	Str  string `json:"str"`
	User *User  `json:"user"`
}
