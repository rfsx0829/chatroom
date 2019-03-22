package service

import (
	"crypto/md5"
	"net/http"

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

func (p *Platform) SignInWithName(w http.ResponseWriter, r *http.Request, name, pass string) (int, error) {
	return p.signIn(w, r, name, "", pass)
}

func (p *Platform) SignInWithEmail(w http.ResponseWriter, r *http.Request, email, pass string) (int, error) {
	return p.signIn(w, r, "", email, pass)
}

func (p *Platform) signIn(w http.ResponseWriter, r *http.Request, name, email, pass string) (int, error) {
	info, err := mysql.Default.GetUserInfo(name, email)
	if err != nil {
		return 0, err
	}

	if info.Pass != Hash(pass) {
		return 0, InvalidPass
	}

	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return 0, err
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

func Hash(source string) string {
	h := md5.New()
	h.Write([]byte(source))
	return string(h.Sum(nil))
}
