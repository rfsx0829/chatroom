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
	err := json.NewDecoder(r.Body).Decode(&x)
	if err != nil {
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
		log.Println(err)
		return
	}

	c.plat.AddConn(conn)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// DelUser delete user
func (c *Controller) DelUser(w http.ResponseWriter, r *http.Request) {
	var x struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&x)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("[DU]", x)

	c.plat.DelUser(x.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Create room
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var x struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&x)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("[CR]", x)

	c.plat.CreateRoom(x.Name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// DelRoom delete room
func (c *Controller) DelRoom(w http.ResponseWriter, r *http.Request) {
	var x struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&x)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("[DR]", x)

	c.plat.DeleteRoom(x.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Enter room
func (c *Controller) Enter(w http.ResponseWriter, r *http.Request) {
	var x struct {
		UID int `json:"uid"`
		RID int `json:"rid"`
	}
	err := json.NewDecoder(r.Body).Decode(&x)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("[ER]", x)

	c.plat.Enter(x.UID, x.RID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Leave room
func (c *Controller) Leave(w http.ResponseWriter, r *http.Request) {
	var x struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&x)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("[LR]", x)

	c.plat.Leave(x.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
