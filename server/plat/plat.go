package plat

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/rfsx0829/chatroom/server/redis"
)

// Platform platform
type Platform struct {
	RoomTable map[int]*Room
	UserTable map[int]*User
	ConnPool  map[int]*websocket.Conn

	waitConn []int

	database redis.Client
}

// New create a platform
func New(db redis.Client) *Platform {
	p := &Platform{
		RoomTable: make(map[int]*Room),
		UserTable: make(map[int]*User),
		ConnPool:  make(map[int]*websocket.Conn),
		waitConn:  make([]int, 0, 2),
		database:  db,
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

func (p *Platform) routine(id int, conn *websocket.Conn) {
	var x wsMes

	for {
		err := conn.ReadJSON(&x)
		if err != nil {
			log.Println("[routine] Error. Will Delete Conn.", err)
			p.UserTable[id].inWhichRoom.removeUser(id)
			p.UserTable[id].inWhichRoom = nil
			p.ConnPool[id] = nil
			return
		}

		log.Println("[routine]", x)

		p.broadCastMes(&x)
	}
}

func (p *Platform) getUnusedID() int {
	for i := 2; i < 300; i++ {
		if _, ok := p.RoomTable[i]; !ok {
			return i
		}
	}

	return -1
}
