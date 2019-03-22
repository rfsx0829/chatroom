package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rfsx0829/chatroom/server/service"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	log.Println("Controller.CreateRoom...")

	if checkMethod(w, r, "POST") {
		log.Println("NotPost")
		return
	}

	var x struct {
		Uid   int    `json:"uid"`
		Rname string `json:"rname"`
		Rpass string `json:"rpass"`
		Pri   bool   `json:"pri"`
	}

	err := unmarshalBody(r, &x)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	log.Println(x)

	rm := service.Client.CreateRoom(x.Rname, x.Rpass, x.Pri, x.Uid)
	data, _ := json.Marshal(struct {
		Rid int `json:"rid"`
	}{rm.Id})

	w.WriteHeader(http.StatusOK)
	w.Write(data)
	log.Println("Create Room id=", rm.Id)
}

func EnterRoom(w http.ResponseWriter, r *http.Request) {
	log.Println("Controller.EnterRoom...")

	if checkMethod(w, r, "POST") {
		log.Println("NotPost")
		return
	}

	var x struct {
		Uid   int    `json:"uid"`
		Rid   int    `json:"rid"`
		Rpass string `json:"rpass"`
	}

	err := unmarshalBody(r, &x)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	log.Println(x)

	err = service.Client.EnterRoom(x.Uid, x.Rid, x.Rpass)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK !"))
	log.Println("Enter Room OK !")
}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	log.Println("Controller.LeaveRoom...")

	if checkMethod(w, r, "POST") {
		log.Println("NotPost")
		return
	}

	var x struct {
		Uid int `json:"uid"`
	}

	err := unmarshalBody(r, &x)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	log.Println(x)

	err = service.Client.LeaveRoom(x.Uid)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK !"))
	log.Println("Leave Room OK !")
}
