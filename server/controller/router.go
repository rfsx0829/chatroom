package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Handler(w http.ResponseWriter, r *http.Request) {
	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	go func(conn *websocket.Conn) {
		var (
			x   FormData
			res = Response{
				Status: http.StatusBadRequest,
			}
		)

		for {
			err := conn.ReadJSON(&x)
			if err != nil {
				res.Text = err.Error()
				con.WriteJSON(res)
			}

			if x.Oper == Close {
				res.Status = http.StatusOK
				res.Text = "Close Connection !"
				con.WriteJSON(res)
				con.Close()
			}

			if dl, found := ops[x.Oper]; found {
				err = dl(&x, &res)
				if err != nil {
					res.Text = err.Error()
					con.WriteJSON(res)
				} else {
					res.Status = http.StatusOK
					con.WriteJSON(res)
				}
			}
		}
	}(con)
}
