package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rfsx0829/chatroom/server/plat"
)

// AddUser temp
func (c *Controller) AddUser(w http.ResponseWriter, r *http.Request) {
	var x plat.Message

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[AU]", x)

	c.plat.AddUser(x.ID, x.Str)

	x.User = &plat.User{
		ID:     x.ID,
		Name:   x.Str,
		Email:  "rfsx0829@163.com",
		Avatar: "http://39.98.162.91:9573/files/picture/a5b23209cbbf66a0ff209e37b37f79d9.png",
	}

	data, err := json.Marshal(x)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// AddConn addconn
func (c *Controller) AddConn(w http.ResponseWriter, r *http.Request) {
	log.Println("[AC]")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	c.plat.AddConn(conn)
}

// DelUser delete user
func (c *Controller) DelUser(w http.ResponseWriter, r *http.Request) {
	var x plat.Message

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[DU]", x)

	c.plat.DelUser(x.User.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
