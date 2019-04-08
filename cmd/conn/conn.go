package conn

import (
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/rfsx0829/chatroom/cmd/types"
	"github.com/rfsx0829/chatroom/cmd/types/oper"
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

func (con Conn) WriteAndRead(fd *types.FormData) (*types.Response, error) {
	var (
		x   types.Response
		err error
	)

	err = con.cn.WriteJSON(fd)
	if err != nil {
		return nil, err
	}

	err = con.cn.ReadJSON(&x)
	return &x, err
}

func (con Conn) Close(fd *types.FormData) error {
	fd.Oper = oper.Close
	return con.cn.WriteJSON(fd)
}