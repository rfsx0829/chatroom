package service

import "time"

func (p *Platform) SendMes(uid int, mes string) error {
	ur := p.getUserById(uid)
	if ur == nil {
		return InvalidUid
	}

	rm := p.getRoomById(ur.RoomId)
	if rm == nil {
		return InvalidRid
	}

	m := &Message{
		Text: mes,
		From: uid,
		To:   -1,
		Time: time.Now(),
	}
	return rm.broadCastMes(m)
}

func (p *Platform) SendToBox(uid, to int, mes string) error {
	ur := p.getUserById(uid)
	if ur == nil {
		return InvalidRid
	}

	tur := p.getUserById(uid)
	if tur == nil {
		return InvalidRid
	}

	mess := Message{
		Text: mes,
		From: uid,
		To:   to,
		Time: time.Now(),
	}

	tur.MessageBox = append(tur.MessageBox, &mess)
	return nil
}

func (r *Room) broadCastMes(mes *Message) (err error) {
	r.MessageQueue = append(r.MessageQueue, mes)

	for _, e := range r.PersonList {
		err = e.Conn.WriteJSON(mes)
		if err != nil {
			return err
		}
	}
	return nil
}
