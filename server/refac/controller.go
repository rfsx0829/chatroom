package refac

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Controller controller
type Controller struct {
	plat *Platform
}

var upgrader websocket.Upgrader

func init() {
	upgrader = websocket.Upgrader{}
}

// NewController create a controller
func NewController(pf *Platform) *Controller {
	return &Controller{pf}
}

// AddUser temp
func (c *Controller) AddUser(w http.ResponseWriter, r *http.Request) {
	var x Message

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[AU]", x)

	c.plat.AddUser(x.ID, x.Str)

	x.User = &User{x.ID, x.Str, "rfsx0829@163.com", "http://39.98.162.91:9573/files/picture/a5b23209cbbf66a0ff209e37b37f79d9.png", nil}
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
	var x Message

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

// CreateRoom create room
func (c *Controller) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var x Message

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[CR]", x)

	c.plat.createRoom(x.Str)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// DeleteRoom delete room
func (c *Controller) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	var x Message

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[DR]", x)

	c.plat.deleteRoom(x.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// EnterRoom enter room
func (c *Controller) EnterRoom(w http.ResponseWriter, r *http.Request) {
	var x Message

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[ER]", x)

	c.plat.enter(x.User.ID, x.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// LeaveRoom leave room
func (c *Controller) LeaveRoom(w http.ResponseWriter, r *http.Request) {
	var x Message

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[LR]", x)

	c.plat.leave(x.User.ID)
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
