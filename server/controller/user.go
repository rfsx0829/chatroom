package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

// AddUser temp
func (c *Controller) AddUser(w http.ResponseWriter, r *http.Request) {
	var x struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
	}

	data, err := func() ([]byte, error) {
		err := json.NewDecoder(r.Body).Decode(&x)
		if err != nil {
			return nil, err
		}

		log.Println("[AU]", x)

		mp, err := c.plat.AddUser(x.Name, x.Pass)
		if err != nil {
			return nil, err
		}

		return json.Marshal(mp)
	}()

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
	var x struct {
		ID int `json:"uid"`
	}

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[DU]", x)

	c.plat.DelUser(x.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
