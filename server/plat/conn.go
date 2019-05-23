package plat

import "github.com/gorilla/websocket"

// AddConn add conn
func (p *Platform) AddConn(conn *websocket.Conn) {
	// TODO: AddConn
	/*if len(p.UserTemp) == 0 {
		return
	}

	id := p.UserTemp[0]
	p.UserTemp = p.UserTemp[1:]

	if u, ok := p.UserTable[id]; ok {
		p.ConnPool[id] = conn
		u.inWhichRoom = p.RoomTable[1]
		p.RoomTable[1].inRoom = append(p.RoomTable[1].inRoom, u)
	}

	go p.routine(id, conn)*/
}
