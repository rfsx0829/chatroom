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
	var x struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("[AU]", x)

	c.plat.AddUser(x.ID, x.Name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
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
		ID int `json:"id"`
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
