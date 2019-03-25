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

func (p *Platform) SendMes(uid int, mes string) error {
	ur := p.GetUserById(uid)
	if ur == nil {
		return InvalidUid
	}

	rm := p.GetRoomById(ur.RoomId)
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
	ur := p.GetUserById(uid)
	if ur == nil {
		return InvalidRid
	}

	tur := p.GetUserById(uid)
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
