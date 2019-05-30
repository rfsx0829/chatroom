package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

// CreateRoom create room
func (c *Controller) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var x struct {
		Name string `json:"name"`
		Pass string `json:"password"`
	}

	if err := func() error {
		err := json.NewDecoder(r.Body).Decode(&x)
		if err != nil {
			return err
		}

		return c.plat.CreateRoom(x.Name, x.Pass)
	}(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[CR]", x)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// DeleteRoom delete room
func (c *Controller) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	var x struct {
		ID int `json:"rid"`
	}

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
	var (
		x struct {
			UID  int    `json:"uid"`
			RID  int    `json:"rid"`
			Pass string `json:"pass"`
		}
		data []byte
	)

	if err := func() error {
		err := json.NewDecoder(r.Body).Decode(&x)
		if err != nil {
			return err
		}

		err = c.plat.Enter(x.UID, x.RID, x.Pass)
		if err != nil {
			return err
		}

		data, err = c.plat.GetRoomMessages(x.RID)
		return err
	}(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[ER]", x)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// LeaveRoom leave room
func (c *Controller) LeaveRoom(w http.ResponseWriter, r *http.Request) {
	var x struct {
		ID int `json:"uid"`
	}

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[LR]", x)

	c.plat.Leave(x.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// RoomList roomlist
func (c *Controller) RoomList(w http.ResponseWriter, r *http.Request) {
	log.Println("[GR]")

	data, err := c.plat.GetRoomList()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
