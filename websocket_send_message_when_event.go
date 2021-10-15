package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var c = make(chan int)

var i int

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/ws", Ws)

	log.Info("Application is running on port 8080..")

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}
	log.Fatal(srv.ListenAndServe())
}

func Home(w http.ResponseWriter, r *http.Request) {
	c <- i
	i++
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Ws(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	fmt.Println("CONNECRED WITH WS")

	go func() {
		for {
			s1 := <-c
			fmt.Println("I HAVE RECEIVED,", s1)
			err := ws.WriteMessage(1, []byte("I AM SENDING A MESSAGE"))
			if err != nil {
				log.Println("write:", err)
				ws.Close()
			}
		}
	}()
}

