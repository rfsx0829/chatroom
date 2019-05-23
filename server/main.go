package main

import (
	"log"
	"net/http"

	"github.com/rfsx0829/chatroom/server/controller"
	"github.com/rfsx0829/chatroom/server/plat"
	"go.uber.org/dig"
)

func main() {
	con := dig.New()

	con.Provide(plat.New)
	con.Provide(controller.New)

	if err := con.Invoke(EntryPoint); err != nil {
		log.Println(err)
	}
}

// EntryPoint run the server
func EntryPoint(c *controller.Controller) error {
	http.HandleFunc("/au", c.AddUser)
	http.HandleFunc("/ac", c.AddConn)
	http.HandleFunc("/du", c.DelUser)
	http.HandleFunc("/cr", c.CreateRoom)
	http.HandleFunc("/dr", c.DeleteRoom)
	http.HandleFunc("/er", c.EnterRoom)
	http.HandleFunc("/lr", c.LeaveRoom)
	http.HandleFunc("/gr", c.RoomList)

	log.Println("Listening At localhost:8089")
	return http.ListenAndServe(":8089", nil)
}
