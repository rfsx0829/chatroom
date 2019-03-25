package types

import (
	"net/url"

	"github.com/gorilla/websocket"
)

type Conn struct {
	cn *websocket.Conn
}

func NewConn(u *url.URL) (Conn, error) {
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return Conn{nil}, err
	}
	return Conn{c}, nil
}

func (con Conn) WriteAndRead(fd *FormData) (*Response, error) {
	var (
		x   Response
		err error
	)

	err = con.cn.WriteJSON(fd)
	if err != nil {
		return nil, err
	}

	err = con.cn.ReadJSON(&x)
	return &x, err
}

func (con Conn) Close(fd *FormData) error {
	fd.Oper = Close
	return con.cn.WriteJSON(fd)
}
