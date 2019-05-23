package plat

import "log"

// Message message
type Message struct {
	ID   int    `json:"id"`
	Str  string `json:"str"`
	User *User  `json:"user"`
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
