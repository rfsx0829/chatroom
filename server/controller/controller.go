package controller

import (
	"github.com/gorilla/websocket"
	"github.com/rfsx0829/chatroom/server/plat"
)

// Controller controller
type Controller struct {
	plat *plat.Platform
}

var upgrader websocket.Upgrader

func init() {
	upgrader = websocket.Upgrader{}
}

// New create a controller
func New(pf *plat.Platform) *Controller {
	return &Controller{pf}
}
