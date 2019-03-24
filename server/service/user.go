package service

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gorilla/websocket"

	"github.com/rfsx0829/chatroom/server/mysql"
)

func (p *Platform) SignUp(name, pass string) error {
	ok, err := mysql.Default.CheckUserName(name)
	if err != nil {
		return err
	}

	if ok {
		return NameExists
	}

	pf := mysql.Profile{
		Name:  name,
		Pass:  Hash(pass),
		Email: "",
	}

	return mysql.Default.AddUser(&pf)
}

func (p *Platform) SignInWithName(con *websocket.Conn, name, pass string) (int, error) {
	return p.signIn(con, name, "", pass)
}

func (p *Platform) SignInWithEmail(con *websocket.Conn, email, pass string) (int, error) {
	return p.signIn(con, "", email, pass)
}

func (p *Platform) signIn(con *websocket.Conn, name, email, pass string) (int, error) {
	info, err := mysql.Default.GetUserInfo(name, email)
	if err != nil {
		return 0, err
	}

	if info.Pass != Hash(pass) {
		return 0, InvalidPass
	}

	user := Person{
		Info:       info,
		Online:     true,
		MessageBox: make([]*Message, 0, 2),
		Conn:       con,
		RoomId:     -1,
		RoomToken:  "",
	}

	p.OnlineList = append(p.OnlineList, &user)
	return info.Uid, nil
}

func (p *Platform) AddEmail(uid int, email string) error {
	return mysql.Default.AddEmail(uid, email)
}

func Hash(source string) string {
	h := md5.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}
