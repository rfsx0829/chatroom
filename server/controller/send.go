package controller

import (
	"log"
	"net/http"

	"github.com/rfsx0829/chatroom/server/service"
)

func SendMes(w http.ResponseWriter, r *http.Request) {
	log.Println("Controller.SendMes...")

	if checkMethod(w, r, "POST") {
		log.Println("NotPost")
		return
	}

	var x struct {
		Uid int    `json:"uid"`
		Mes string `json:"mes"`
	}

	err := unmarshalBody(r, &x)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	log.Println(x)

	err = service.Client.SendMes(x.Uid, x.Mes)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK !"))
	log.Println("Send Mes OK !")
}

func SendToBox(w http.ResponseWriter, r *http.Request) {
	log.Println("Controller.SendToBox...")

	if checkMethod(w, r, "POST") {
		log.Println("NotPost")
		return
	}

	var x struct {
		Uid    int    `json:"uid"`
		SendTo int    `json:"sendto"`
		Mes    string `json:"mes"`
	}

	err := unmarshalBody(r, &x)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	log.Println(x)

	err = service.Client.SendToBox(x.Uid, x.SendTo, x.Mes)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK !"))
	log.Println("Send To Box OK !")
}
