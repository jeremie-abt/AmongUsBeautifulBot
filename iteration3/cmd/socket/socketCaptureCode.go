package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	//	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

import "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/services"

//import "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/framework"
//import "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain"

type socketCaptureCode struct {
	socketAdapter services.ISocketEventAdapter
}

func newSocketCaptureCode(socAdapter services.ISocketEventAdapter) *socketCaptureCode {
	return &socketCaptureCode{
		socketAdapter: socAdapter,
	}
}

var wsConns map[*websocket.Conn]bool = make(map[*websocket.Conn]bool)

const FAKEGAMEID = "asdlkjasdlasd65464"

var upgrader = websocket.Upgrader{} // use default options

var addr = flag.String("addr", "localhost:3000", "http service address")

func handleLobby(w http.ResponseWriter, r *http.Request, msg []byte) {
	socketAdapter := services.NewSocketAdapter()

	socketAdapter.HandleLobbyEvent("asd", "asdasda")
}

func handleState(w http.ResponseWriter, r *http.Request, msg []byte) {

}

func handlePlayer(w http.ResponseWriter, r *http.Request, msg []byte) {

}

func handleConnectCode(w http.ResponseWriter, r *http.Request, msg []byte) {

}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	captureCodeWsEndPoints := map[string]func(
		http.ResponseWriter, *http.Request, []byte){
		"/lobby":       handleLobby,
		"/state":       handleLobby,
		"/player":      handleLobby,
		"/connectCode": handleLobby,
	}

	if err != nil {
		fmt.Printf("upgrader error : %s\n", err)
	}
	defer c.Close()
	for {
		_, buffer, err := c.ReadMessage()
		if err != nil {
			fmt.Printf("Err reading msg : %s\n", err)
		}
		f, found := captureCodeWsEndPoints[r.URL.String()]
		if !found {
			fmt.Printf(
				"cet endpoint n'existe pas ou n'est pas gere : %s", r.URL)
		}
		f(w, r, buffer)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", wsHandler)
	http.HandleFunc("/other", wsHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))

	//fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	//sc := make(chan os.Signal, 1)
	//defer close(sc)
	//signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	//<-sc
}
