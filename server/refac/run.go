package refac

import (
	"log"
	"net/http"
)

// EntryPoint run the server
func EntryPoint(c *Controller) error {
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
