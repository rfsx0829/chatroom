package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("In Handler ...")

	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	log.Println("Go !!!")

	go func(conn *websocket.Conn) {
		var (
			x   FormData
			res = Response{
				Status: http.StatusBadRequest,
			}
		)

		for {
			log.Println("In Loop ...")

			err := conn.ReadJSON(&x)
			if err != nil {
				log.Println(err)

				res.Text = err.Error()
				con.WriteJSON(res)
				continue
			}

			log.Println(x.Oper)

			if x.Oper == Close {
				res.Status = http.StatusOK
				res.Text = "Close Connection !"

				log.Println(res.Text)

				con.WriteJSON(res)
				con.Close()
				return
			}

			if x.Oper == SignIn {
				err = SignInDealer(con, &x, &res)
				if err != nil {
					log.Println(err)

					res.Text = err.Error()
					con.WriteJSON(res)
					con.Close()
					return
				} else {
					log.Println("OK", res.Extra)

					res.Status = http.StatusOK
					con.WriteJSON(res)
				}
				continue
			}

			if dl, found := ops[x.Oper]; found {
				err = dl(&x, &res)
				if err != nil {
					log.Println(err)

					res.Text = err.Error()
					con.WriteJSON(res)
				} else {
					log.Println("OK", res.Extra)

					res.Status = http.StatusOK
					con.WriteJSON(res)
				}
			}
		}
	}(con)
}
