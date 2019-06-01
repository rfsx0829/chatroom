package plat

import "github.com/gorilla/websocket"

// AddConn add conn
func (p *Platform) AddConn(conn *websocket.Conn) {
	// TODO: AddConn
	if len(p.waitConn) == 0 {
		return
	}

	id := p.waitConn[0]
	p.waitConn = p.waitConn[1:]

	if _, ok := p.UserTable[id]; ok {
		p.ConnPool[id] = conn
	}

	go p.routine(id, conn)
}
