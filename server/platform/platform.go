package platform

import (
	"github.com/gorilla/websocket"
)

// Platform expose API to controller
type Platform struct {
	conPool map[int]*websocket.Conn
	online  map[int]*user
}

// DefaultPlatform TODO:
var DefaultPlatform *Platform

func newPlatform() *Platform {
	return &Platform{
		conPool: make(map[int]*websocket.Conn),
	}
}

func initDefault() {
	if DefaultPlatform == nil {
		DefaultPlatform = newPlatform()
	}
}
