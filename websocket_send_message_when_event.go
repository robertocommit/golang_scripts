package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var wsS []string
var c = make(chan string)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{param}", Home)
	router.HandleFunc("/ws/{username}", Ws)
	log.Info("Application is running on port 8080..")
	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}
	log.Fatal(srv.ListenAndServe())
}

func Home(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["param"]
	c <- param
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Ws(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	wsS = append(wsS, username)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	go func() {
		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				return
			}
		}
	}()
	go func() {
		for {
			s1 := <-c
			if s1 == username {
				err := ws.WriteMessage(1, []byte("I AM SENDING USER DATA"))
				if err != nil {
					ws.Close()
					return
				}
			}
		}
	}()
}
