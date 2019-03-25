package service

import (
	"time"
)

type Message struct {
	Text string    `json:"text"`
	From int       `json:"from"`
	To   int       `json:"to"`
	Time time.Time `json:"time"`
}

func (p *Platform) SendMes(uid int, mes *Message) error {
	ur := p.GetUserById(uid)
	if ur == nil {
		return InvalidUid
	}

	rm := p.GetRoomById(ur.RoomId)
	if rm == nil {
		return InvalidRid
	}

	return rm.broadCastMes(mes)
}

func (p *Platform) SendToBox(uid, to int, mes *Message) error {
	ur := p.GetUserById(uid)
	if ur == nil {
		return InvalidRid
	}

	tur := p.GetUserById(uid)
	if tur == nil {
		return InvalidRid
	}

	tur.MessageBox = append(tur.MessageBox, mes)
	return nil
}
