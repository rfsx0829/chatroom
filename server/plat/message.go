package plat

import "log"

type wsMes struct {
	UID     int    `json:"uid"`
	Type    int    `json:"type"` // 0 for text, 1 for image
	Content string `json:"content"`
}

func (p *Platform) broadCastMes(mes *wsMes) {
	if u, ok := p.UserTable[mes.UID]; ok && u.inWhichRoom != nil {
		u.inWhichRoom.messages = append(u.inWhichRoom.messages, mes)
		for _, e := range u.inWhichRoom.inRoom {
			p.sendToConn(e.ID, mes)
		}
	}
}

func (p *Platform) sendToConn(to int, mes *wsMes) {
	if con, ok := p.ConnPool[to]; ok && con != nil {
		if err := con.WriteJSON(mes); err != nil {
			log.Printf("[%d, %v]: %s", to, mes, err.Error())
		}
	}
}
