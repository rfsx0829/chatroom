package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rfsx0829/chatroom/server/plat"
)

// CreateRoom create room
func (c *Controller) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var x plat.Message

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[CR]", x)

	c.plat.CreateRoom(x.Str)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// DeleteRoom delete room
func (c *Controller) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	var x plat.Message

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[DR]", x)

	c.plat.DeleteRoom(x.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// EnterRoom enter room
func (c *Controller) EnterRoom(w http.ResponseWriter, r *http.Request) {
	var x plat.Message

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[ER]", x)

	c.plat.Enter(x.User.ID, x.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// LeaveRoom leave room
func (c *Controller) LeaveRoom(w http.ResponseWriter, r *http.Request) {
	var x plat.Message

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[LR]", x)

	c.plat.Leave(x.User.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// RoomList roomlist
func (c *Controller) RoomList(w http.ResponseWriter, r *http.Request) {
	log.Println("[GR]")

	list := c.plat.GetRoomList()

	data, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
